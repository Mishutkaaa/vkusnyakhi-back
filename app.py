from flask import Flask, jsonify, request
import psycopg2
import os
from dotenv import load_dotenv

app = Flask(__name__)

@app.after_request
def after_request(response):
    response.headers.add('Access-Control-Allow-Origin', '*')
    response.headers.add('Access-Control-Allow-Headers', 'Content-Type,Authorization')
    response.headers.add('Access-Control-Allow-Methods', 'GET,PUT,POST,DELETE,OPTIONS')
    return response

load_dotenv()

database = os.environ.get("database")
user = os.environ.get("user")
password = os.environ.get("password")

def get_db():
    return psycopg2.connect(
        host='localhost',
        database=database, 
        user=user,
        password=password
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

@app.route('/categories')
def get_categories():
    category_type = request.args.get('type') 
    
    conn = get_db()
    cur = conn.cursor()

    query = "SELECT id, name FROM public.categoriess"
    params = []

    if category_type and category_type in ['drinks', 'food']:
        query += " WHERE type = %s::product_type"
        params.append(category_type)
    
    cur.execute(query, params)
    rows = cur.fetchall() 

    categories = []
    for row in rows:
        categories.append({
            "id": row[0],
            "name": row[1]
        })
    
    cur.close()
    conn.close()
    
    return jsonify(categories)

@app.route('/newProduct', methods=['POST'])
def add_product():
    data = request.get_json()  
    
    conn = get_db()
    cur = conn.cursor()

    cur.execute(f"SELECT MAX(id) FROM public.{data['db']}")
    max_id = cur.fetchone()[0] or 0
    new_id = max_id + 1
    
    table_name = 'public.' + data['db']

    categories = data.get('categories', [])
    categories_str = "{" + ",".join(map(str, categories)) + "}" if categories else "{}"
    
    cur.execute(
        f"INSERT INTO {table_name} (id, name, image, categories) VALUES (%s, %s, %s, %s)",
        (new_id, data['name'], data['img'], categories_str)
    )
    
    conn.commit()  
    cur.close()
    conn.close()
    
    return jsonify({"status": "ok", "message": "ok"})

if __name__ == '__main__':
    app.run(debug=True)