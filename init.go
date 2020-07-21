package main

func init() {
	if err := setupLogger(); err != nil {
		panic("Fail to setup logger:" + err.Error())
	}
}
