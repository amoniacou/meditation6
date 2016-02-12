#!/bin/bash
mongo persons --eval "db.dropDatabase();"
time bundle exec ruby ./populator.rb
ab -c 100 -n 10000 localhost:8080/persons
