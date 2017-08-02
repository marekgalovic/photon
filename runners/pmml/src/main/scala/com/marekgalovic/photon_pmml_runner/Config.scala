package com.marekgalovic.photon_pmml_runner

// import org.rogach.scallop._

class Config(args: Seq[String]) {
  val address = "0.0.0.0"
  val port = 5022
  val modelsDir = "./"
}
