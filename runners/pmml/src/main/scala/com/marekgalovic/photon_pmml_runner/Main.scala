package com.marekgalovic.photon_pmml_runner

import java.io.File

object Main {
  private val evaluator = new Evaluator

  def main(args: Array[String]) {
    val conf = new Config(args)

    println(conf.modelsDir)
    println(conf.nodeIp)

    modelFiles(conf.modelsDir).foreach { println }
  }

  private def modelFiles(dir: String): List[File] = {
    val d = new File(dir)
    if (!d.exists) {
      throw new Exception("Models dir does not exists")
    }
    if (!d.isDirectory) {
      throw new Exception("Models dir is not a directory")
    }
    return d.listFiles.filter(f => f.isFile && f.getName.endsWith("xml")).toList 
  }
}
