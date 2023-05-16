# Chat:

### Функциональность:
1. Возможность создания пользователем чата с другим сотрудником агентства;
2. Возможность создания пользователем группового чата с другим сотрудниками агентства;
3. Просмотр пользователем истории сообщений в чатах;
4. Возможность отправки и получения текстовых сообщений в индивидуальных и групповых чатах с коллегами;
5. Отображение статуса сообщений: отправлено/прочитано;
6. Отображение последней активности собеседника в чате;
7. Возможность отправки и получения сообщений, содержащих медиа-файлы;
8. Возможность редактирования и удаления сообщения из чата;
9. Возможность создания тредов и обсуждения в них необходимых вопросов несколькими участниками переписки.
10. Поиск
12. Отправка и получение голосовых сообщений;

### Конфигурация:
* Установка зависимостей: 
```
go mod tidy
```
* Кодогенерация, если были изменены protobuf api
``` 
make generate
```
* Инициализация БД, если требуется локальная
```
make init-local-db
```
* Настроки окружения указываются в файле .env

### Запуск:
```
make run
```