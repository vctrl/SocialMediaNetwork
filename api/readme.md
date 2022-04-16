
// USERS

    GET /me
    response: user

    POST /register
    request: { login, password, name, surname, age, sex, interests, city }
    response: user, 201
    
    POST /login
    
    request: { login, password }
    response: {user_id, token}
    
    GET /users/{id1,id2,id3...}

    PUT /users/{id}
    response user
    
    DELETE /users/{id}
    response: 200 ok

    POST /logout
// FRIENDS

    // список друзей
    GET /friends/
    response {[ user ]}
    
    // исходящие заявки
    GET /friends/requests/sent

    response [user_id1, user_id2 ...] 

    // входящие заявки
    GET /friends/requests

    response [user_id1, user_id2 ...] 

    // отправить заявку в друзья
    POST /friends/{user_id}
    
    response 200

    // принять заявку в друзья
    POST /friends/{id}/accept

    response 200

    // удалить друга
    DELETE /friends/{user_id}

    // отменить заявку
    DELETE /friends/requests/{id}


user {id, login, name, surname, age, sex, interests, city }

friend_request {}