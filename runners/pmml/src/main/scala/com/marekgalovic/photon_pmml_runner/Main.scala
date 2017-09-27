package com.marekgalovic.photon_pmml_runner

import io.grpc.{Server, ServerBuilder}
import photon.runner.RunnerServiceGrpc
import scala.concurrent.ExecutionContext

import org.apache.curator.framework.{CuratorFramework, CuratorFrameworkFactory}
import org.apache.curator.retry.ExponentialBackoffRetry

object Main {
  def main(args: Array[String]) {
    val config = new Config(args)
    val zookeeper = getZookeeperClient(config)
    val modelManager = new ModelManager(config, zookeeper)
    modelManager.load

    val server = ServerBuilder.forPort(config.port())
      .addService(RunnerServiceGrpc.bindService(new RunnerService(modelManager), ExecutionContext.global))
      .build
      .start

    println("Models dir: "+config.modelsDir())
    println("gRPC server listening on: "+config.port())

    sys.addShutdownHook {
      zookeeper.close
      server.shutdown
    }
    
    server.awaitTermination
  }

  private def getZookeeperClient(config: Config): CuratorFramework = {
    val client = CuratorFrameworkFactory.builder
      .connectString(config.zookeeperNodes())
      .retryPolicy(new ExponentialBackoffRetry(1000, 3))
      .namespace(config.zookeeperBasepath())
      .build

    client.start
    client.getZookeeperClient.blockUntilConnectedOrTimedOut

    client
  }
}
