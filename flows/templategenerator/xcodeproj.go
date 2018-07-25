package templategenerator

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"os"
)

type AddedFile struct {
	Path     string
	Filename string
}
type XcodeProjTargetAddedFiles struct {
	TargetName string
	AddedFiles []AddedFile
}
type XcodeProjAdded struct {
	XcodeprojFilePath string
	TargetFiles       []XcodeProjTargetAddedFiles
}

func xcodeproj(addedTargetFiles []XcodeProjAdded) {
	templateString := templateRubyFile()
	t := template.Must(template.New("ruby").Parse(templateString))

	file, err := os.OpenFile("xcode.rb", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	err = t.Execute(writer, addedTargetFiles)
	if err != nil {
		log.Println("executing template:", err)
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}

	out, err := execCmd("chmod", "+x", "xcode.rb")
	if err != nil {
		panic(err)
	}
	out, err = execCmd("./xcode.rb")
	if err != nil {
		panic(err)
	}
	fmt.Printf("in all caps: %s\n", out)

	os.Remove("xcode.rb")
}

func templateRubyFile() string {
	return `#!/usr/bin/env ruby

require 'xcodeproj'

class AddedFile
	attr_accessor :path_folder, :file_name

	def initialize(path_folder, file_name)
		@path_folder = path_folder
		@file_name = file_name
	end
end

def obtain_target(project, target_name)
	project.targets.each do |target|
		return target if target.name == target_name
	end
	return nil
end

def split_path(path)
	path.to_s.split('/')
end

def new_group(project, path)
	group = project
	group_names = split_path(path)

	group_names.each do |group_name|
		next_group = group[group_name]

		unless next_group
			next_group = group.new_group(group_name, group_name)
		end

		group = next_group
	end

	group
end

def group_files_contains(group, file_name)
	group.files.each do |file|
		if file.path == file_name then
			return true
		end
	end

	return false
end

def add_files_to_target(project, target_name, files)
	target = obtain_target(project, target_name)
	unless target
		put 'Error target ' + target_name + ' not found'
		return
	end
	
	
	files.each do |file|
		group = new_group(project, file.path_folder)
		isContains = group_files_contains(group, file.file_name)
	
		ref_added_file = group.new_file(file.file_name) unless isContains
		target.add_file_references([ref_added_file]) unless isContains
	end
end

{{range .}}
	project_path = '{{.XcodeprojFilePath}}'
	project = Xcodeproj::Project.open(project_path)

	{{range .TargetFiles}}
		target_name = '{{.TargetName}}'
		files = [
			{{range .AddedFiles}}
			AddedFile.new('{{.Path}}', '{{.Filename}}'),
			{{end}}
			]

		add_files_to_target(project, target_name, files)
	{{end}}
	
	project.save
{{end}}`
}
