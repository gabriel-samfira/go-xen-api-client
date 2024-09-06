//
// This file is generated. To change the content of this file, please do not
// apply the change to this file because it will get overwritten. Instead,
// change xenapi.go and execute 'go generate'.
//

package xenapi

import (
	"fmt"
	"github.com/amfranz/go-xmlrpc-client"
	"reflect"
	"strconv"
	"time"
)

var _ = fmt.Errorf
var _ = xmlrpc.NewClient
var _ = reflect.TypeOf
var _ = strconv.Atoi
var _ = time.UTC

type EventOperation string

const (
	// An object has been created
	EventOperationAdd EventOperation = "add"
	// An object has been deleted
	EventOperationDel EventOperation = "del"
	// An object has been modified
	EventOperationMod EventOperation = "mod"
)

type EventRecord struct {
	// An ID, monotonically increasing, and local to the current session
	ID int64
	// The time at which the event occurred
	Timestamp time.Time
	// The name of the class of the object that changed
	Class string
	// The operation that was performed
	Operation EventOperation
	// A reference to the object that changed
	Ref string
	// The uuid of the object that changed
	ObjUUID string
}

type EventRef string

// Asynchronous event registration and handling
type EventClass struct {
	client *Client
}

// Inject Injects an artificial event on the given object and return the corresponding ID
func (_class EventClass) Inject(sessionID SessionRef, class string, ref string) (_retval string, _err error) {
	_method := "event.inject"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_classArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "class"), class)
	if _err != nil {
		return
	}
	_refArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "ref"), ref)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _classArg, _refArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// GetCurrentID Return the ID of the next event to be generated by the system
func (_class EventClass) GetCurrentID(sessionID SessionRef) (_retval int64, _err error) {
	_method := "event.get_current_id"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

// From Blocking call which returns a new token and a (possibly empty) batch of events. The returned token can be used in subsequent calls to this function.
//
// Errors:
//  SESSION_NOT_REGISTERED - This session is not registered to receive events.  You must call event.register before event.next.  The session handle you are using is echoed.
//  EVENTS_LOST - Some events have been lost from the queue and cannot be retrieved.
func (_class EventClass) From(sessionID SessionRef, classes []string, token string, timeout float64) (_retval []EventRecord, _err error) {
	_method := "event.from"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_classesArg, _err := convertStringSetToXen(fmt.Sprintf("%s(%s)", _method, "classes"), classes)
	if _err != nil {
		return
	}
	_tokenArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "token"), token)
	if _err != nil {
		return
	}
	_timeoutArg, _err := convertFloatToXen(fmt.Sprintf("%s(%s)", _method, "timeout"), timeout)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _classesArg, _tokenArg, _timeoutArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEventRecordSetToGo(_method + " -> ", _result.Value)
	return
}

// Next Blocking call which returns a (possibly empty) batch of events. This method is only recommended for legacy use. New development should use event.from which supercedes this method. 
//
// Errors:
//  SESSION_NOT_REGISTERED - This session is not registered to receive events.  You must call event.register before event.next.  The session handle you are using is echoed.
//  EVENTS_LOST - Some events have been lost from the queue and cannot be retrieved.
func (_class EventClass) Next(sessionID SessionRef) (_retval []EventRecord, _err error) {
	_method := "event.next"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEventRecordSetToGo(_method + " -> ", _result.Value)
	return
}

// Unregister Unregisters this session with the event system
func (_class EventClass) Unregister(sessionID SessionRef, classes []string) (_err error) {
	_method := "event.unregister"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_classesArg, _err := convertStringSetToXen(fmt.Sprintf("%s(%s)", _method, "classes"), classes)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _classesArg)
	return
}

// Register Registers this session with the event system.  Specifying * as the desired class will register for all classes.
func (_class EventClass) Register(sessionID SessionRef, classes []string) (_err error) {
	_method := "event.register"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_classesArg, _err := convertStringSetToXen(fmt.Sprintf("%s(%s)", _method, "classes"), classes)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _classesArg)
	return
}
