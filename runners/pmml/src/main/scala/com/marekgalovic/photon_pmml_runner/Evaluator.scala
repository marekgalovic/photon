package com.marekgalovic.photon_pmml_runner

import scala.collection.mutable.{Map => MutableMap}
import org.dmg.pmml.PMML
import org.jpmml.evaluator.{Evaluator => JPmmlEvaluator, ModelEvaluatorFactory}

class Evaluator {
  private var models: MutableMap[String,JPmmlEvaluator] = MutableMap[String,JPmmlEvaluator]()
  private val modelsEvaluatorFactory = ModelEvaluatorFactory.newInstance

  def addModel(uid: String, pmml: PMML) = {
    models.put(uid, modelsEvaluatorFactory.newModelEvaluator(pmml))
  }

  def removeModel(uid: String) = {
    models.remove(uid)
  }
}
