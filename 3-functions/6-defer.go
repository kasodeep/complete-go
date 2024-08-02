func GetUsername(dstName, srcName string) (user string, err error) {
	// Open a connection to a database
	conn, _ := db.Open(srcName)
 
	// Close the source file *anywhere* the GetUsername function returns.
	defer conn.Close()
 
	username, err := db.FetchUser()
	if err != nil{
	  return 0, err
	}	
	return username, nil
 }