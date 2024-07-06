#SET SQL_SAFE_UPDATES = 0;
#DELETE FROM kmkglass.products;
#DELETE FROM kmkglass.typeproduct;
#DELETE FROM kmkglass.optionsproduct;


import mysql.connector
from faker import Faker
import time
from PIL import Image, ImageDraw
import io



car_models = [
    "Camry", "Accord", "Civic", "Corolla", "Mustang",
    "F-150", "Silverado", "Ram 1500", "Model 3", "Model S",
    "Rav4", "CR-V", "CX-5", "Outback", "Forester",
    "Altima", "Sentra", "Maxima", "Pathfinder", "Rogue",
    "Explorer", "Escape", "Edge", "Tahoe", "Suburban",
    "Equinox", "Malibu", "Impala", "Sierra", "Yukon",
    "Wrangler", "Cherokee", "Grand Cherokee", "Compass", "Renegade",
    "Charger", "Challenger", "Durango", "Fusion", "Focus",
    "Tucson", "Elantra", "Sonata", "Santa Fe", "Palisade",
    "Optima", "Sorento", "Sportage", "Stinger", "Telluride",
    "3 Series", "5 Series", "7 Series", "X3", "X5",
    "C-Class", "E-Class", "S-Class", "GLC", "GLE",
    "A3", "A4", "A6", "Q5", "Q7",
    "Golf", "Passat", "Tiguan", "Jetta", "Beetle",
    "Impreza", "Legacy", "BRZ", "911", "Cayenne",
    "Macan", "Panamera", "Boxster", "Range Rover", "Discovery",
    "Defender", "F-Type", "XF", "XC60", "XC90",
    "S60", "S90", "V60", "Chiron", "Veyron",
    "Continental", "Flying Spur", "Ghost", "Phantom", "Wraith",
    "DB11", "Vantage", "Huracan", "Aventador", "LaFerrari"
]

car_brands = [
    "Toyota", "Ford", "Chevrolet", "Honda", "Nissan",
    "BMW", "Mercedes-Benz", "Volkswagen", "Audi", "Hyundai",
    "Kia", "Lexus", "Subaru", "Mazda", "Buick",
    "Cadillac", "Chrysler", "Dodge", "GMC", "Jeep",
    "Ram", "Tesla", "Volvo", "Jaguar", "Land Rover",
    "Mitsubishi", "Mini", "Porsche", "Infiniti", "Acura",
    "Alfa Romeo", "Fiat", "Genesis", "Lincoln", "Maserati",
    "Bentley", "Ferrari", "Lamborghini", "Rolls-Royce", "Aston Martin",
    "Bugatti", "Pagani", "Koenigsegg", "Rivian", "Lucid",
    "Peugeot", "Renault", "Citroën", "Skoda", "Seat",
    "Opel", "Saab", "Lancia", "Holden", "Vauxhall",
    "Dacia", "Geely", "Great Wall", "Chery", "BYD",
    "Haval", "Tata", "Mahindra", "Maruti Suzuki", "Proton",
    "Perodua", "SsangYong", "Daewoo", "Isuzu", "Suzuki",
    "Smart", "Scion", "Saturn", "Oldsmobile", "Plymouth",
    "Pontiac", "Mercury", "Hummer", "Fisker", "Spyker",
    "Rover", "MG", "Austin", "Triumph", "Alpina",
    "Ariel", "Caterham", "Donkervoort", "Gumpert", "TVR",
    "Wiesmann", "Zenos", "Zenvo", "DeLorean", "Lotus",
    "Morgan", "Noble", "BAC", "Radical", "Vector"
]

num_rows_products = 1000000
num_rows_typeproducts = 2000
num_rows_optionsproducts = 2000
number_of_entries_for_the_commit = 1000

# Настройте подключение к базе данных MySQL
connection = mysql.connector.connect(
    host="localhost",
    user="root",
    password="123437",
    database="kmkglass"
)
cursor = connection.cursor()

# Создайте объект Faker
fake = Faker()

# Функция для генерации случайного изображения
def generate_image():
    width, height = 320, 200
    image = Image.new('RGB', (width, height), (255, 255, 255))
    draw = ImageDraw.Draw(image)
    
    # Генерация случайного текста и рисование на изображении
    text = fake.text(max_nb_chars=20)
    draw.text((10, 10), text, fill=(0, 0, 0))
    
    # Сохранение изображения в бинарный формат
    img_byte_arr = io.BytesIO()
    image.save(img_byte_arr, format='PNG')
    img_byte_arr = img_byte_arr.getvalue()
    return img_byte_arr

def add_table_products(num_rows):
    # Укажите имя таблицы и её колонки
    table_name = "products"
    columns = "(idproducts, price, name, year, article, brand, model, length, photo, width, inStock, amount, idTypeProduct, idOptions)"
    # Начинайте вставку данных
    for i in range(num_rows):
        idproducts = i+1
        price = fake.random_int(1, 99999)
        name = fake.name()
        year = fake.date_of_birth(minimum_age=18, maximum_age=90)
        article = fake.random_int(1, 99999)
        brand = fake.sentence(nb_words = 1, ext_word_list=car_brands)
        model = fake.sentence(nb_words = 1, ext_word_list=car_models)
        length = fake.random_int(1, 99999)
        photo = generate_image()
        width = fake.random_int(1, 99999)
        inStock = fake.boolean()
        amount = fake.random_int(1, 99999)
        idTypeProduct = fake.random_int(1, num_rows_typeproducts)
        idOptions = fake.random_int(1, num_rows_optionsproducts)

        # Создайте SQL-запрос для вставки данных
        insert_query = f"""
        INSERT INTO {table_name} {columns} 
        VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)
        """
        cursor.execute(insert_query, (idproducts, price, name, year, article, brand, model, length, photo, width, inStock, amount, idTypeProduct, idOptions))

        # Периодически выполняйте коммит, чтобы не перегружать память
        if i % number_of_entries_for_the_commit == 0:
            connection.commit()
            print(f"Inserted {i} rows...")

    # Финальный коммит
    connection.commit()

def add_table_typeProducts(num_rows):
    # Укажите имя таблицы и её колонки
    table_name = "typeproduct"
    columns = "(idTypeProduct, name)"

    # Начинайте вставку данных
    for i in range(num_rows):
        idTypeProduct = i+1
        name = fake.name()

        # Создайте SQL-запрос для вставки данных
        insert_query = f"""
        INSERT INTO {table_name} {columns} 
        VALUES (%s, %s)
        """
        cursor.execute(insert_query, (idTypeProduct, name))

        # Периодически выполняйте коммит, чтобы не перегружать память
        if i % number_of_entries_for_the_commit == 0:
            connection.commit()
            print(f"Inserted {i} rows...")

    # Финальный коммит
    connection.commit()
    
def add_table_optionsproduct(num_rows):
    # Укажите имя таблицы и её колонки
    table_name = "optionsproduct"
    columns = "(idoptions, name)"

    # Начинайте вставку данных
    for i in range(num_rows):
        idoptions = i+1
        name = fake.name()

        # Создайте SQL-запрос для вставки данных
        insert_query = f"""
        INSERT INTO {table_name} {columns} 
        VALUES (%s, %s)
        """
        cursor.execute(insert_query, (idoptions, name))

        # Периодически выполняйте коммит, чтобы не перегружать память
        if i % number_of_entries_for_the_commit == 0:
            connection.commit()
            print(f"Inserted {i} rows...")    


start_time = time.time()
add_table_typeProducts(num_rows_typeproducts)
add_table_optionsproduct(num_rows_optionsproducts)
add_table_products(num_rows_products)


# Закройте соединение
cursor.close()
connection.close()

end_time = time.time()
print(f"{end_time - start_time} seconds.")