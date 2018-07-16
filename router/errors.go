package router

import "errors"

var errNoArguments = errors.New("Аргументы не найдены")
var errHelpNoArguments = errors.New("jessica help <command>")
var errRouteNotFound = errors.New("Команда не найдена")
