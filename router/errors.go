package router

import "errors"

var errHelpNoArguments = errors.New("jessica help <command>")
var errRouteNotFound = errors.New("Команда не найдена")
