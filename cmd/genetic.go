/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"giali.com/commander/cmd/interfaces"
	"giali.com/commander/cmd/util"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// geneticCmd represents the genetic command
var geneticCmd = &cobra.Command{
	Use:   "genetic",
	Short: "cli docker-compose.yaml generator",
	Long:  `Generatore di docker-compose tramite cli`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("genetic called")
		handler(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(geneticCmd)

	// Here you will define your flags and configuration settings.
	geneticCmd.Flags().BoolP("nosql", "n", false, "add db to docker-compose")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// geneticCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// geneticCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func handler(cmd *cobra.Command, args []string) {
	_, err := cmd.Flags().GetBool("nosql")
	if err != nil {
		fmt.Println(err)
	}

	service_map := make(map[string]interfaces.DockerService)

	if util.SelectYN("Vuoi aggiungere mongodb?") {
		service_map["mongodb"] = buildMongoDB()
	}
	if util.SelectYN("Vuoi aggiungere postgres?") {
		service_map["postgres"] = buildPostgres()
	}
	if util.SelectYN("Vuoi aggiungere neo4j?") {
		service_map["neo4j"] = buildNeo4j()
	}
	if util.SelectYN("Vuoi aggiungere redis?") {
		service_map["redis"] = buildRedis()
	}
	if util.SelectYN("Vuoi aggiungere rabbitmq?") {
		service_map["rabbitmq"] = buildRabbitMQ()
	}
	if util.SelectYN("Vuoi aggiungere kafka,zookeeper e kafka_ui?") {
		service_map["zookeeper"] = buildZookeeper()
		service_map["kafka"] = buildKafka()
		service_map["kafka_ui"] = buildKafkaUI()
	}

	docker_compose := interfaces.DockerCompose{
		Version:  "3",
		Services: service_map,
	}

	yamlData, _ := yaml.Marshal(&docker_compose)
	util.CreateFile("docker-compose.yml", string(yamlData))

	if util.SelectYN("Vuoi lanciare i container?") {
		execDockerCompose()
	}

}
func buildMongoDB() interfaces.DockerService {
	return interfaces.DockerService{
		Restart:       "always",
		Image:         "mongo",
		ContainerName: "mongodb",
		Volumes:       []string{"./init-mongo.js:/docker-entrypoint-initdb.d/init-mongo.js:ro", "./mongo-volume:/data/database"},
		Ports:         []string{"27017-27019:27017-27019"},
	}
}
func buildNeo4j() interfaces.DockerService {
	return interfaces.DockerService{
		Restart:       "unless-stopped",
		Image:         "neo4j",
		ContainerName: "neo4j",
		Ports:         []string{"7474:7474", "7687:7687"},
		Volumes: []string{
			"./conf:/conf",
			"./data:/data",
			"./import:/import",
			"./logs:/logs",
			"./plugins:/plugins",
		},
		Environment: []string{
			"NEO4J_dbms_memory_pagecache_size=1G",
			"NEO4J_dbms.memory.heap.initial_size=1G",
			"NEO4J_dbms_memory_heap_max__size=1G",
		},
	}
}
func buildRedis() interfaces.DockerService {
	return interfaces.DockerService{
		Restart:       "unless-stopped",
		Image:         "redis",
		ContainerName: "redis",
		Ports:         []string{"6379:6379"},
		Expose:        []string{"6379"},
	}
}
func buildRabbitMQ() interfaces.DockerService {
	return interfaces.DockerService{
		Restart:     "always",
		Image:       "rabbitmq",
		Environment: []string{"RABBITMQ_DEFAULT_USER=guest", "RABBITMQ_DEFAULT_PASS=guest"},
		Ports:       []string{"5672:5672", "15672:15672"},
	}
}
func buildZookeeper() interfaces.DockerService {
	return interfaces.DockerService{
		ContainerName: "zookeeper",
		Image:         "confluentinc/cp-zookeeper:latest",
		Environment:   []string{"ZOOKEEPER_CLIENT_PORT=2181", "ZOOKEEPER_TICK_TIME=2000"},
		Ports:         []string{"22181:2181"},
	}
}
func buildKafka() interfaces.DockerService {
	return interfaces.DockerService{
		ContainerName: "kafka",
		Image:         "confluentinc/cp-kafka:latest",
		DependsOn:     []string{"zookeeper"},
		Environment: []string{
			"KAFKA_BROKER_ID=1",
			"KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181",
			"KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://kafka:9092,PLAINTEXT_HOST://localhost:29092",
			"KAFKA_LISTENER_SECURITY_PROTOCOL_MAP=PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT",
			"KAFKA_INTER_BROKER_LISTENER_NAME=PLAINTEXT",
			"KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1",
		},
		Ports: []string{"29092:29092"},
	}
}
func buildKafkaUI() interfaces.DockerService {
	return interfaces.DockerService{
		ContainerName: "kafka_ui",
		Image:         "provectuslabs/kafka-ui:latest",
		DependsOn:     []string{"kafka"},
		Environment: []string{
			"KAFKA_CLUSTERS_0_ZOOKEEPER=zookeeper:2181",
			"KAFKA_CLUSTERS_0_NAME=local",
			"KAFKA_CLUSTERS_0_BOOTSTRAPSERVERS=kafka:9092",
		},
		Ports: []string{"8080:8080"},
	}
}

func buildPostgres() interfaces.DockerService {
	return interfaces.DockerService{
		ContainerName: "postgres",
		Image:         "postgres:latest",
		Restart:       "always",
		Environment: []string{
			"POSTGRES_USER=postgres",
			"POSTGRES_PASSWORD=postgres",
		},
		Volumes: []string{"/var/lib/postgresql/data"},
		Ports:   []string{"5432:5432"},
	}
}
func execDockerCompose() {
	command := exec.Command("docker-compose", "up", "-d", "--build")
	stdout, err := command.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := command.Start(); err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(stdout)

	if err != nil {
		log.Fatal(err)
	}

	if err := command.Wait(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", string(data))
}
