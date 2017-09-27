package com.marekgalovic.photon_pmml_runner

import java.net.InetAddress
import org.rogach.scallop._

class Config(args: Seq[String]) extends ScallopConf(args) {
  val modelsDir = opt[String](required=true, default=Option(sys.env.getOrElse("PHOTON_MODELS_DIR", "./models")))
  val port = opt[Int](required=true, default=Option(5006))
  val zookeeperNodes = opt[String](required=true, default=Option("127.0.0.1:2181"), name="zookeeper.nodes")
  val zookeeperBasepath = opt[String](required=true, default=Option("photon"), name="zookeeper.basepath")
  verify()

  def nodeIp: String = {
    return InetAddress.getLocalHost.getHostAddress
  }
}
