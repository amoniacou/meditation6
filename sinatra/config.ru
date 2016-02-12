require 'mongoid'
require 'sinatra/base'
require 'sinatra/json'
require_relative 'server/model'
require_relative 'server/server'
Mongoid.load!("./mongoid.yml", :development)
run Meditation::Server
