package com.marekgalovic.photon_pmml_runner

import scala.collection.mutable.{Map => MutableMap}

import java.io.{File, FileInputStream}
import java.nio.file.{Path, Paths}
import java.nio.file.StandardWatchEventKinds.{ENTRY_CREATE, ENTRY_DELETE}

import akka.actor.ActorSystem
import com.beachape.filemanagement.MonitorActor
import com.beachape.filemanagement.RegistryTypes.Callback
import com.beachape.filemanagement.Messages.RegisterCallback

import org.dmg.pmml.PMML
import org.xml.sax.InputSource
import org.jpmml.model.{ImportFilter, JAXBUtil}
import org.jpmml.evaluator.{Evaluator, ModelEvaluatorFactory}

import org.apache.curator.framework.CuratorFramework
import org.apache.zookeeper.CreateMode

class ModelManager(config: Config, zk: CuratorFramework) {
  private val modelsDirMonitorActor = ActorSystem("actorSystem").actorOf(MonitorActor(concurrency = 2))
  modelsDirMonitorActor ! RegisterCallback(event = ENTRY_CREATE, path = Paths.get(config.modelsDir()), callback = createModel)
  modelsDirMonitorActor ! RegisterCallback(event = ENTRY_DELETE, path = Paths.get(config.modelsDir()), callback = deleteModel)

  private val modelEvaluatorFactory = ModelEvaluatorFactory.newInstance

  private var models: MutableMap[String,Evaluator] = MutableMap[String,Evaluator]()
  private var zookeeperUids: MutableMap[String,String] = MutableMap[String,String]()

  def get(versionUid: String): Evaluator = {
    models(versionUid)
  }

  def load = {
    val dir = new File(config.modelsDir())
    if (!dir.exists) {
      throw new Exception("Models dir does not exists.")
    }
    if (!dir.isDirectory) {
      throw new Exception("Models dir is not a directory.")
    }

    dir.listFiles.foreach { file =>
      createModel(file.toPath)
    }
  }

  private def createModel(path: Path) = {
    val fileNameParts = path.getFileName.toString.split("\\.")
    if (!(fileNameParts.length == 2 && fileNameParts.last == "xml")) {
      throw new Exception("Invalid file name: "+path.getFileName)
    }

    try {
      val pmml = readPmml(path)
      val evaluator = modelEvaluatorFactory.newModelEvaluator(pmml)

      registerVersionToZookeeper(fileNameParts(0))
      models.put(fileNameParts(0), evaluator)
      println("Deployed model version: "+fileNameParts(0))
    } catch {
      case e: Exception => println("Failed to deploy model: "+fileNameParts(0), e)
    }
  }

  private def deleteModel(path: Path) = {
    val fileNameParts = path.getFileName.toString.split("\\.")
    if (!(fileNameParts.length == 2 && fileNameParts.last == "xml")) {
      throw new Exception("Invalid file name: "+path.getFileName)
    }

    unregisterVersionFromZookeeper(fileNameParts(0))
    models.remove(fileNameParts(0))
    println("Deleted model version: "+fileNameParts(0))
  }

  private def readPmml(path: Path): PMML = {
    val inputStream = new FileInputStream(path.toFile)
    try {
      val pmmlSource = ImportFilter.apply(new InputSource(inputStream))
      return JAXBUtil.unmarshalPMML(pmmlSource)
    } finally {
      inputStream.close()  
    }
  }

  def registerVersionToZookeeper(versionUid: String) = {
    val nodeInfo = ("{\"address\":\""+config.nodeIp+"\", \"port\":"+config.port+"}").getBytes("UTF-8")

    val createdPath = zk.create
      .creatingParentsIfNeeded
      .withProtection
      .withMode(CreateMode.EPHEMERAL_SEQUENTIAL)
      .forPath("/instances/"+versionUid+"/r-", nodeInfo)

    zookeeperUids.put(versionUid, createdPath.split("/").last)
  }

  def unregisterVersionFromZookeeper(versionUid: String) = {
    if (!zookeeperUids.contains(versionUid)) {
      throw new Exception("Zookeeper uid not found for model version: "+versionUid)
    }

    zk.delete
      .forPath("/instances/"+versionUid+"/"+zookeeperUids(versionUid))

    zookeeperUids.remove(versionUid)
  }
}
