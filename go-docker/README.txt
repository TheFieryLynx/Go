1.  Для начала необходимо построить Dockerfile, если это не сделано
    (Если он уже построен, переходим к п.2)
    В терминале нужно войти в папку со всей программой и написать команду:
    docker build -t go-docker .

2.  Запуск Docker:
    docker run -d -p 8080:8080 go-docker
    Docker присвоился CONTAINER ID

    Для просмотра всех ID работающих контейнеров:
    docker container ls

3.  Для взаимодействия с сервером можно прописывать команду:
    curl http://localhost:8080?name=smb

4.  docker container stop CONTAINER ID


#Если нет разрешений для команды docker, то нужно писать sudo перед каждой командой.