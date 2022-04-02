
// USERS
POST /users/register
    request: { login, password, name, surname, age, sex, interests, city }
    response: user, 201
    
    POST /users/login
    
    request: { login, password }
    response: user
    
    GET /users/{id1,id2,id3...}

    PUT /users/{id}
    response user
    
    DELETE /users/{id}
    response: 200 ok

// FRIENDS

    // список друзей
    GET /friends/
    response {[ user ]}
    
    // исходящие заявки
    GET /friends/sent_requests

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


user {id, login, name, surname, age, sex, interests, city }

friend_request {}