#!/usr/bin/env bash
if [ $# -ne 1 ]; then
  echo "need param" 1>&2
  exit 1
fi

mainFile=./src/$1.go
testFile=./src/$1_test.go

cat << EOS > $mainFile
package exercise

func $1 {
}

EOS

echo "gen src: ${mainFile}"

cat << EOS > $testFile
package exercise

import (
	"testing"
)

func Test$1(t *testing.T) {
}

EOS

echo "gen test: ${testFile}"
