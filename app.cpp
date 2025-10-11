#include "crow.h"
#include <iostream>

int main() {
    crow::SimpleApp app;

    CROW_ROUTE(app, "/")([](){
        return "Hello from C++ Backend! Server is working!";
    });

    CROW_ROUTE(app, "/api/hello")([](){
        crow::json::wvalue response;
        response["message"] = "Hello World!";
        response["status"] = "success";
        return response;
    });

    CROW_ROUTE(app, "/api/users")
    ([](){
        std::vector<crow::json::wvalue> users;
        
        users.push_back(crow::json::wvalue{
            {"id", 1},
            {"name", "John Doe"}
        });
        
        users.push_back(crow::json::wvalue{
            {"id", 2},
            {"name", "Jane Smith"}
        });

        crow::json::wvalue response;
        response["users"] = std::move(users);
        return response;
    });

    CROW_ROUTE(app, "/health")([](){
        return "Server is healthy!";
    });

    std::cout << "=== C++ Backend Server ===" << std::endl;
    std::cout << "Server running on: http://localhost:8080" << std::endl;
    std::cout << "Available endpoints:" << std::endl;
    std::cout << "  GET /" << std::endl;
    std::cout << "  GET /api/hello" << std::endl;
    std::cout << "  GET /api/users" << std::endl;
    std::cout << "  GET /health" << std::endl;
    std::cout << "Press Ctrl+C to stop the server" << std::endl;

    app.port(8080).multithreaded().run();
    
    return 0;
}