# https://nightlies.apache.org/flink/flink-docs-master/api/python/reference/pyflink.datastream/index.html

from pyflink.common.typeinfo import Types
from pyflink.datastream import StreamExecutionEnvironment
from pyflink.datastream.connectors.kafka import FlinkKafkaConsumer
from pyflink.datastream.formats.json import JsonRowDeserializationSchema

env = StreamExecutionEnvironment.get_execution_environment()

env.add_jars(
    "/Users/rushiyadwade/Documents/go_dir/source/iudx-task/task_runners/flink-sql-connector-kafka.jar"
)

deserialization_schema = (
    JsonRowDeserializationSchema.builder()
    .type_info(type_info=Types.ROW([Types.INT(), Types.STRING()]))
    .build()
)

kafka_consumer = FlinkKafkaConsumer(
    topics="data-kaveri",
    deserialization_schema=deserialization_schema,
    properties={"bootstrap.servers": "localhost:9092", "group.id": "flink-job"},
)

ds = env.add_source(kafka_consumer)

# references to check more on specifying a sink for the flink job
# https://stackoverflow.com/q/68620906
ds.add_sink()
