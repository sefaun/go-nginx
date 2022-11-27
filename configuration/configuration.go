package configuration

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Configuration struct{}

func (configuration *Configuration) ReadConfiguration() ([]string, []string) {
	f, err := os.Open("default.conf")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	content := bufio.NewScanner(f)
	var line_counter = 0
	var server_hosts = []string{}
	var stream_hosts = []string{}

	for content.Scan() {
		if line_counter == 0 {
			server_hosts = configuration.seperateServerPortsScope(content.Text())
		}
		if line_counter == 1 {
			stream_hosts = configuration.seperateStreamPortsScope(content.Text())
		}
		line_counter++
	}

	if err := content.Err(); err != nil {
		log.Fatal(err)
	}

	return server_hosts, stream_hosts
}

func (_ *Configuration) seperateServerPortsScope(content string) []string {
	first_scope := strings.Index(string(content), "[") + 1
	last_scope := strings.Index(string(content), "]")

	server_urls := content[first_scope:last_scope]

	commad_index := strings.Index(string(content), ",")
	//Just One Server
	if commad_index == -1 {
		server_urls = strings.Trim(server_urls, " ")
		return []string{server_urls}
	}

	servers := strings.Split(content[first_scope:last_scope], ",")

	for i := 0; i < len(servers); i++ {
		servers[i] = strings.Trim(servers[i], " ")
	}

	return servers
}

func (_ *Configuration) seperateStreamPortsScope(content string) []string {
	first_scope := strings.Index(string(content), "[") + 1
	last_scope := strings.Index(string(content), "]")

	stream_urls := content[first_scope:last_scope]

	commad_index := strings.Index(string(content), ",")
	//Just One Server
	if commad_index == -1 {
		stream_urls = strings.Trim(stream_urls, " ")
		return []string{stream_urls}
	}

	streams := strings.Split(content[first_scope:last_scope], ",")

	for i := 0; i < len(streams); i++ {
		streams[i] = strings.Trim(streams[i], " ")
	}

	return streams
}
