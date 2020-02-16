package constants

const (
	StorageFor = "JOIN users ON users.schedule_storage_id = schedule_storages.id AND users.telegram_chat_id = ?"
)
