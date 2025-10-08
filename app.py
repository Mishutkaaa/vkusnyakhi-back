from flask import Flask, jsonify
import psycopg2
app = Flask(__name__)

@app.after_request
def after_request(response):
    response.headers.add('Access-Control-Allow-Origin', '*')
    response.headers.add('Access-Control-Allow-Headers', 'Content-Type,Authorization')
    response.headers.add('Access-Control-Allow-Methods', 'GET,PUT,POST,DELETE,OPTIONS')
    return response
def get_db():
    return psycopg2.connect(
        host='localhost',
        database='mishutka', 
        user='mishutka',
        password='mishutka'
    )

@app.route('/')
def home():
    return jsonify({"status": "ok", "message": "Hello"})

@app.route('/drinks')
def get_drinks():
    conn = get_db()
    cur = conn.cursor()
    
    cur.execute("SELECT id, name, image FROM public.drinks")
    rows = cur.fetchall() 

    drinks = []
    for row in rows:
        drinks.append({
            "id": row[0],
            "name": row[1],
            "img": row[2]
        })
    
    cur.close()
    conn.close()
    
    return jsonify(drinks)

@app.route('/food')
def get_food():
    conn = get_db()
    cur = conn.cursor()
    
    cur.execute("SELECT id, name, image FROM public.food")
    rows = cur.fetchall() 

    food = []
    for row in rows:
        food.append({
            "id": row[0],
            "name": row[1],
            "img": row[2]
        })
    
    cur.close()
    conn.close()
    
    return jsonify(food)



if __name__ == '__main__':
    app.run(debug=True)