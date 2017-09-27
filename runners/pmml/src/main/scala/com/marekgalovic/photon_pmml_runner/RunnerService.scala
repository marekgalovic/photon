package com.marekgalovic.photon_pmml_runner

import photon.runner.{RunnerServiceGrpc, RunnerEvaluateRequest, RunnerEvaluateResponse}
import photon.core.ValueInterface
import scala.concurrent.Future
import scala.collection.JavaConversions._

import java.util.LinkedHashMap
import org.dmg.pmml.FieldName
import org.jpmml.evaluator.{Evaluator, InputField}


class RunnerService(modelManager: ModelManager) extends RunnerServiceGrpc.RunnerService {
  def evaluate(req: RunnerEvaluateRequest) = {
    try {
      val evaluator = modelManager.get(req.versionUid)
      val result = evaluator.evaluate(protoBufToInputValues(evaluator, req.features))
      Future.successful(RunnerEvaluateResponse(result=resultToProtoBuf(evaluator, result.asInstanceOf[LinkedHashMap[FieldName,Any]])))
    } catch {
      case e: Exception => Future.failed(e)
    }
  }

  private def protoBufToInputValues(evaluator: Evaluator, features: Map[String,ValueInterface]): Map[FieldName,Any] = {
    evaluator.getActiveFields.map{ field => 
      val value = features(field.getName.toString).value
      val resultValue = value match {
        case ValueInterface.Value.ValueBoolean(v) => v
        case ValueInterface.Value.ValueInt32(v) => v
        case ValueInterface.Value.ListInt32(v) => v
        case ValueInterface.Value.ValueInt64(v) => v
        case ValueInterface.Value.ListInt64(v) => v
        case ValueInterface.Value.ValueFloat32(v) => v
        case ValueInterface.Value.ListFloat32(v) => v
        case ValueInterface.Value.ValueFloat64(v) => v
        case ValueInterface.Value.ListFloat64(v) => v
        case ValueInterface.Value.ValueString(v) => v
        case ValueInterface.Value.ValueBytes(v) => v
        case ValueInterface.Value.Empty => null
      }
      (field.getName -> resultValue)
    }.toMap
  }

  private def resultToProtoBuf(evaluator: Evaluator, result: LinkedHashMap[FieldName,Any]): Map[String,ValueInterface] = {
    evaluator.getOutputFields.map { field =>
      val value = result(field.getName)
      val valueInterfaceProto = value match {
        case v: Boolean => ValueInterface().withValueBoolean(v)
        case v: Int => ValueInterface().withValueInt32(v)
        case v: Long => ValueInterface().withValueInt64(v)
        case v: Float => ValueInterface().withValueFloat32(v)
        case v: Double => ValueInterface().withValueFloat64(v)
        case v: String => ValueInterface().withValueString(v)
        case _ => throw new Exception("Invalid value type")
      }
      (field.getName.toString -> valueInterfaceProto)
    }.toMap
  }
}
