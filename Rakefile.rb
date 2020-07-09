require 'pathname'
require 'rake/clean'
require 'tmpdir'

task default: "generate_v2_go_protos"

protos  = Rake::FileList.new('test_protos/*.proto').map { |p| p.split('/').last }

task :generate_v2_go_protos do
  path = './test_protos'
  
  protos.each do |t|
    sh "docker run -v $(pwd):/defs namely/protoc-all -i #{path} -f #{t} -l go -o #{path}"
  end
end

task :generate_v1_go_protos do
  in_path = './test_protos'
  
  protos.each do |t|
    path = 'test_protos'
    sh "docker run -v $(pwd):/defs namely/protoc-all:1.29_1 -i #{in_path} -f #{t} -l go -o #{out_path}"
  end
end

