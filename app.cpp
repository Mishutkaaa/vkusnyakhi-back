#include "crow.h"
#include <iostream>

int main() {
    crow::SimpleApp app;

    CROW_ROUTE(app, "/drinks")([](){
crow::json::wvalue x({{"id", "row[0]"}}); //Select from drinks where row = 1
    x["name"] = "row[1]";
    x["image"] = "row[2]";
    return x;
    });
    CROW_ROUTE(app, "/food")([](){
      crow::json::wvalue x({{"id", "row[0]"}});
    x["name"] = "row[1]";
    x["image"] = "row[2]";
    return x;
    });

    CROW_ROUTE(app, "/categories")([](){
        crow::json::wvalue response;
        // METHOD GET
        // РУЧКА ДЛЯ ПОЛУЧАНИЯ КАТЕГОРИЙ, ОТДАЁМ ОТВЕТ В ВИДЕ 
        // {
        // "id":1, 
        // "name":"name"
        // }
        return response;
    });

    CROW_ROUTE(app, "/newProduct")([](){
        crow::json::wvalue response;
        // METHOD POST
        // РУЧКА ДЛЯ СОЗДАНИЯ СУЩНОСТИ, ПРИНИМАЕМ ДАННЫЕ В ВИДЕ
        // {
        //  name: "имя_продукта",
        // img: "фотка_продукта",
        // db: "имя_бд"
        // }
        return response;
    });

    std::cout << "=== C++ Backend Server ===" << std::endl;
    std::cout << "Server running on: http://localhost:8080" << std::endl;
    std::cout << "Available endpoints:" << std::endl;

    std::cout << "  GET /drinks" << std::endl;
    std::cout << "  GET /food" << std::endl;
    std::cout << "  GET /newProduct" << std::endl;
    std::cout << "Press Ctrl+C to stop the server" << std::endl;

    app.port(8080).multithreaded().run();
    
}