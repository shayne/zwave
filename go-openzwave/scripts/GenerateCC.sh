#!/usr/bin/env bash

PREFIX=CC

enumerate()
{
    grep [iI]mplements /libs/openzwave-1.4.164/cpp/src/command_classes/*.h | tr -d '()' | sed "s|.*/command_classes/||;s|:.*COMMAND_CLASS_| |;s|,.*||;s|.*\\.h ||;s/ x/ 0x/"
}

symbol()
{
    local t=$1
    echo $(echo $t | sed "s/\(.\)\([A-Z]\)/\1_\2/g" | tr [a-z] [A-Z])
}

mkdir -p $PREFIX && cat > $PREFIX/$PREFIX.go <<EOF
package $PREFIX;

//
// *** generated by scripts/$(basename $0)
//

// DO NOT EDIT THIS FILE

import "fmt"

const UNKNOWN = 0xff
var UNKNOWN_ENUM = Enum{UNKNOWN, "$PREFIX.ENUM"}

const (
$(x=0; enumerate | while read code value; do echo "   $code = $value"; let x=x+1; done)
)

var (
    enums = [...]Enum{
$(x=0; enumerate | while read code value; do echo "      { $code, \"$PREFIX.$code\" },"; let x=x+1; done)
		UNKNOWN_ENUM }

    fromName = make(map[string]*Enum)
    fromCode = make(map[int]*Enum)
)

type Enum struct {
     Code int
     Name string
}

func init() {
     for _, e := range enums {
        fromName[e.Name] = &e
        fromCode[e.Code] = &e
     }
}


func ToEnum(code int) *Enum {	
     e,ok := fromCode[code]
     if ok {
        return e
     } else {
        return &UNKNOWN_ENUM
     }
}

func FromName(name string) *Enum {	
     e,ok := fromName[name]
     if ok {
        return e
     } else {
        return &UNKNOWN_ENUM
     }
}

func (val Enum) IsValid() bool {
    return val != UNKNOWN_ENUM
}

func (val Enum) String() string {
     if val.IsValid() {
	return val.Name
     } else { 
        return fmt.Sprintf("%s[%d]", UNKNOWN_ENUM.Name, val.Code);
     }	
}

EOF
gofmt -s -w $PREFIX/$PREFIX.go && cd $PREFIX && go install 

