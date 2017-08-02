name := "pmml_runner"

version := "1.0"

scalaVersion := "2.12.1"

libraryDependencies ++= Seq(
  "org.rogach" %% "scallop" % "3.0.3",
  "org.jpmml" % "pmml-evaluator" % "1.3.7",
  "org.apache.curator" % "curator-framework" % "2.6.0",
  "org.apache.curator" % "curator-recipes" % "2.6.0"
)
