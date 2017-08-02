package com.marekgalovic.photon_pmml_runner

import java.io.File
import org.rogach.scallop._

object Main {
  def main(args: Array[String]) {
    val conf = new Config(args)

    println(conf.modelsDir)
  }
}
