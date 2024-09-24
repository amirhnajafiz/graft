package v1

func (r Router) GetLogs() {
	list := r.Storage.Fetch()
	response := make([]logResponse, len(list))

	for index, item := range list {
		response[index].Term = int(item.GetTerm())
		response[index].Index = int(item.GetIndex())
		response[index].Command = item.GetCommand()
	}
}
