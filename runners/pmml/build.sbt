name := "photon_pmml_runner"
version := "1.0"
description := "Photon PMML runner service."

scalaVersion := "2.12.1"

enablePlugins(JavaAppPackaging)
enablePlugins(DockerPlugin)
enablePlugins(AshScriptPlugin)

import com.typesafe.sbt.packager.docker._
maintainer in Docker := "Marek Galovic <galovic.galovic@gmail.com>"
dockerBaseImage := "openjdk:8-jre"
dockerRepository := Some("marekgalovic")
dockerUpdateLatest := true

libraryDependencies ++= Seq(
  "org.rogach" %% "scallop" % "3.0.3",
  "org.jpmml" % "pmml-evaluator" % "1.3.7",
  "io.grpc" % "grpc-netty" % com.trueaccord.scalapb.compiler.Version.grpcJavaVersion,
  "com.trueaccord.scalapb" %% "scalapb-runtime-grpc" % com.trueaccord.scalapb.compiler.Version.scalapbVersion,
  "com.beachape.filemanagement" %% "schwatcher" % "0.3.2",
  "org.apache.curator" % "curator-framework" % "2.6.0",
  "org.apache.curator" % "curator-recipes" % "2.6.0"
)

assemblyMergeStrategy in assembly := {
 case PathList("META-INF", xs @ _*) => MergeStrategy.discard
 case x => MergeStrategy.first
}

PB.protoSources in Compile := Seq(sourceDirectory.value / "../../../protos")
PB.targets in Compile := Seq(
  scalapb.gen() -> (sourceManaged in Compile).value
)
