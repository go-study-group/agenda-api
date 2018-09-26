package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"

	"github.com/globalsign/mgo"
)

func dialInfo(host, port, dbName, username, pass string) *mgo.DialInfo {
	// DialInfo holds options for establishing a session with a MongoDB cluster.
	addr := fmt.Sprintf("%s:%s", host, port)
	return &mgo.DialInfo{
		Addrs:    []string{addr},
		Timeout:  60 * time.Second,
		Database: dbName,
		Username: username,
		Password: pass,
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
	}
}

// Create a session which maintains a pool of socket connections
// to our Azure Cosmos DB MongoDB database.
//
// make sure to close the session after you're done
func session(dialInfo *mgo.DialInfo) (*mgo.Session, error) {

	sess, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return nil, err
	}
	sess.SetSafe(&mgo.Safe{})
	return sess, nil
}
