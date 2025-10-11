#include "crow.h"
#include <iostream>

int main() {
    crow::SimpleApp app;

    CROW_ROUTE(app, "/")([](){
        return "Всё ок";
    });

    CROW_ROUTE(app, "/drinks")([](){
        crow::json::wvalue response;
        // METHOD GET
        // РУЧКА ДЛЯ ПОЛУЧАНИЯ НАПИТКОВ, ОТДАЁМ ОТВЕТ В ВИДЕ 
        // {
        // "id":1, 
        // "name":"name",
        //  "image":"image"
        // }
        return response;
    });

    CROW_ROUTE(app, "/food")
    ([](){
        crow::json::wvalue response;
        // METHOD GET
        // РУЧКА ДЛЯ ПОЛУЧАНИЯ ЕДЫ, ОТДАЁМ ОТВЕТ В ВИДЕ 
        // {
        // "id":1, 
        // "name":"name",
        //  "image":"image"
        // }
        return response;
    });

    CROW_ROUTE(app, "/newProduct")([](){
        // METHOD POST
        // РУЧКА ДЛЯ СОЗДАНИЯ СУЩНОСТИ, ПРИНИМАЕМ ДАННЫЕ В ВИДЕ
        // {
        //  name: "имя_продукта",
        // img: "фотка_продукта",
        // db: "имя_бд"
        // }
        return
    });

    std::cout << "=== C++ Backend Server ===" << std::endl;
    std::cout << "Server running on: http://localhost:8080" << std::endl;
    std::cout << "Available endpoints:" << std::endl;
    std::cout << "  GET /" << std::endl;
    std::cout << "  GET /drinks" << std::endl;
    std::cout << "  GET /food" << std::endl;
    std::cout << "  GET /newProduct" << std::endl;
    std::cout << "Press Ctrl+C to stop the server" << std::endl;

    app.port(8080).multithreaded().run();
    
}