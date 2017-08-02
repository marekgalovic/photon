package com.marekgalovic.photon_pmml_runner

import java.net.InetAddress
// import org.rogach.scallop._

class Config(args: Seq[String]) {
  val modelUid = ""
  val modelsDir = "./"
  val address = "0.0.0.0"
  val port = 5022

  def nodeIp: String = {
    return InetAddress.getLocalHost.getHostAddress
  }
}
