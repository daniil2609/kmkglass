#SET SQL_SAFE_UPDATES = 0;
#DELETE FROM kmkglass.products;
#DELETE FROM kmkglass.brands;
#DELETE FROM kmkglass.glass_options;
#DELETE FROM kmkglass.glass_types;
#DELETE FROM kmkglass.models;
#DELETE FROM kmkglass.year_model;


import mysql.connector
from faker import Faker
import time
from PIL import Image, ImageDraw
import io
import random
from minio import Minio
from minio.error import S3Error


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

glass_options_list = [
    "Силикатное стекло", "Закаленное стекло", "Ламинированное стекло", "Сапфировое стекло",
    "Оптическое стекло", "Кварцевое стекло", "Боросиликатное стекло", "Флоат-стекло",
    "Огнеупорное стекло", "Ультрафиолетовое стекло", "Инфракрасное стекло", "Цветное стекло",
    "Молочное стекло", "Антибликовое стекло", "Антистатическое стекло", "Гидрофобное стекло",
    "Фотохромное стекло", "УФ-защитное стекло", "Теплозащитное стекло", "Антирефлексное стекло",
    "Антиграффити стекло", "Бронированное стекло", "Акустическое стекло", "Электропроводное стекло",
    "Электрохромное стекло", "Спектрально-избирательное стекло", "Мультифункциональное стекло",
    "Градиентное стекло", "Триплекс стекло", "Зеркальное стекло", "Солнцезащитное стекло",
    "Светоотражающее стекло", "Матированное стекло", "Антимикробное стекло", "Стеклокерамика",
    "Вакуумное стекло", "Горячелитое стекло", "Пирооптическое стекло", "Огнестойкое стекло",
    "Декоративное стекло", "Пескоструйное стекло", "Рифленое стекло", "Узорчатое стекло",
    "Мозаичное стекло", "Тонированное стекло", "Катаное стекло", "Листовое стекло",
    "Стеклопакет", "Фиберглас", "Армированное стекло", "Металлизированное стекло",
    "Нано-стекло", "Фотонное стекло", "Биостекло", "Керамическое стекло", "Витражное стекло",
    "Флоат-стекло с покрытием", "Вакуумно-осажденное стекло", "Стекло с низким содержанием железа",
    "Стекло с высокой прозрачностью", "Ультратонкое стекло", "Ультраплотное стекло",
    "Магнитное стекло", "Ионно-обменное стекло", "Стекло с наночастицами", "Рентгенозащитное стекло",
    "Радиационно-стойкое стекло", "Микропористое стекло", "Пьезоэлектрическое стекло",
    "Фотоэлектрическое стекло", "Теплоизоляционное стекло", "Шумоизоляционное стекло",
    "Стекло с антиконденсатным покрытием", "Стекло с антибактериальным покрытием",
    "Стекло с антиоксидантным покрытием", "Стекло с противоударным покрытием",
    "Стекло с антибликовым покрытием", "Стекло с самоочищающимся покрытием", "Стекло с водоотталкивающим покрытием",
    "Стекло с пылеотталкивающим покрытием", "Стекло с солнцезащитным покрытием",
    "Стекло с инфракрасным покрытием", "Стекло с ультрафиолетовым покрытием",
    "Стекло с антистатическим покрытием", "Стекло с антицарапным покрытием", "Стекло с декоративным покрытием",
    "Стекло с антифоговым покрытием", "Стекло с зеркальным покрытием", "Стекло с прозрачным покрытием",
    "Стекло с цветным покрытием", "Стекло с градиентным покрытием", "Стекло с матовым покрытием",
    "Стекло с узорчатым покрытием", "Стекло с пескоструйным покрытием", "Стекло с рифленым покрытием"
]

auto_glass_types = [
    "Лобовое стекло", 
    "Боковое стекло передней двери", 
    "Боковое стекло задней двери",
    "Стекло заднего окна", 
    "Крыша с панорамным стеклом", 
    "Стекло люка",
    "Лобовое стекло с подогревом",
    "Лобовое стекло с обогревом зоны покоя щеток", 
    "Лобовое стекло с инфракрасной защитой",
    "Лобовое стекло с УФ-защитой", 
    "Лобовое стекло с датчиком дождя", 
    "Лобовое стекло с датчиком света",
    "Лобовое стекло с проекционным дисплеем", 
    "Лобовое стекло с системой антизапотевания",
    "Лобовое стекло с фотохромным покрытием", 
    "Лобовое стекло с антибликовым покрытием",
    "Лобовое стекло с водоотталкивающим покрытием", 
    "Лобовое стекло с встроенной антенной",
    "Лобовое стекло с защитой от шума", 
    "Лобовое стекло с дополнительной шумоизоляцией"
]

year_ranges = [
    "1925-1929", "1930-1934", "1935-1939", "1940-1944", "1945-1949", 
    "1950-1954", "1955-1959", "1960-1964", "1965-1969", "1970-1974", 
    "1975-1979", "1980-1984", "1985-1989", "1990-1994", "1995-1999", 
    "2000-2004", "2005-2009", "2010-2014", "2015-2019", "2020-2024", 
    "1924-1928", "1929-1933", "1934-1938", "1939-1943", "1944-1948", 
    "1949-1953", "1954-1958", "1959-1963", "1964-1968", "1969-1973", 
    "1974-1978", "1979-1983", "1984-1988", "1989-1993", "1994-1998", 
    "1999-2003", "2004-2008", "2009-2013", "2014-2018", "2019-2023", 
    "1923-1927", "1928-1932", "1933-1937", "1938-1942", "1943-1947", 
    "1948-1952", "1953-1957", "1958-1962", "1963-1967", "1968-1972", 
    "1973-1977", "1978-1982", "1983-1987", "1988-1992", "1993-1997", 
    "1998-2002", "2003-2007", "2008-2012", "2013-2017", "2018-2022", 
    "1922-1926", "1927-1931", "1932-1936", "1937-1941", "1942-1946", 
    "1947-1951", "1952-1956", "1957-1961", "1962-1966", "1967-1971", 
    "1972-1976", "1977-1981", "1982-1986", "1987-1991", "1992-1996", 
    "1997-2001", "2002-2006", "2007-2011", "2012-2016", "2017-2021", 
    "1921-1925", "1926-1930", "1931-1935", "1936-1940", "1941-1945", 
    "1946-1950", "1951-1955", "1956-1960", "1961-1965", "1966-1970", 
    "1971-1975", "1976-1980", "1981-1985", "1986-1990", "1991-1995", 
    "1996-2000", "2001-2005", "2006-2010", "2011-2015", "2016-2020"
]

num_rows_products = 1000000
num_rows_typeproducts = len(auto_glass_types)
num_rows_optionsproducts = len(glass_options_list)
num_rows_brands = len(car_brands)
num_rows_models = len(car_models)
num_rows_year_model = len(year_ranges)
number_of_entries_for_the_commit = 10000

# Настройка подключение к базе данных MySQL
connection = mysql.connector.connect(
    host="localhost",
    port=3306,
    user="root",
    password="123437",
    database="kmkglass"
)
cursor = connection.cursor()

# Настройка MinIO клиента
minio_client = Minio(
    'localhost:9000',
    access_key='root',
    secret_key='123437123437',
    secure=False
)

# Убедимся, что ведро существует
bucket_name = 'kmkglass-photo-bucket'
if not minio_client.bucket_exists(bucket_name):
    minio_client.make_bucket(bucket_name)
    
# Создайте объект Faker
fake = Faker()

# Функция для генерации случайного изображения
def generate_image(width, height):
    image = Image.new('RGB', (width, height), (random.randint(0, 254), random.randint(0, 254), random.randint(0, 254)))
    draw = ImageDraw.Draw(image)
    
    # Генерация случайного текста и рисование на изображении
    text = fake.text(max_nb_chars=50)
    draw.text((10, 10), text, fill=(0, 0, 0))
    
    # Сохранение изображения в бинарный формат
    img_byte_arr = io.BytesIO()
    image.save(img_byte_arr, format='PNG')
    img_byte_arr = img_byte_arr.getvalue()
    return img_byte_arr

# Сохранение фотографии в MinIO
def upload_photo_to_minio(photo, file_name):
    try:
        minio_client.put_object(
            bucket_name,
            file_name,
            io.BytesIO(photo), 
            length=len(photo),
            content_type='image/png'
            
        )
    except S3Error as exc:
        print("Error occurred:", exc)

# Получение ссылки для доступа к фотографии
def get_photo_url(file_name):
    try:
        url = minio_client.presigned_get_object(bucket_name, file_name)
        return url
    except S3Error as exc:
        print("Error occurred:", exc)
        return None

def add_table_products(num_rows):
    # Укажите имя таблицы и её колонки
    table_name = "products"
    columns = "(idproducts, price, name, article, brands_name, models_name, year_model_name, length, photo, width, amount, glass_types_name, glass_options_name)"
    
    #вставка URL фото:
    # Генерация и загрузка в minio фотографии
    file_name = "products_"+str(1)+".png"
    upload_photo_to_minio(generate_image(350, 200), file_name)
    # Получение ссылки из minio
    photo_url = get_photo_url(file_name)
    
    # Начинайте вставку данных
    for i in range(num_rows):
        idproducts = i+1
        price = fake.random_int(1, 99999)
        name = fake.name()
        article = fake.random_int(1, 99999)
        brands_name = random.choice(car_brands)
        models_name = random.choice(car_models)
        year_model_name = random.choice(year_ranges)
        length = fake.random_int(1, 99999)
        photo = photo_url
        width = fake.random_int(1, 99999)
        amount = fake.random_int(1, 99999)
        glass_types_name = random.choice(auto_glass_types)
        glass_options_name = random.choice(glass_options_list)
        

        

        # Создайте SQL-запрос для вставки данных
        insert_query = f"""
        INSERT INTO {table_name} {columns} 
        VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)
        """
        cursor.execute(insert_query, (idproducts, price, name, article, brands_name, models_name, year_model_name, length, photo, width, amount, glass_types_name, glass_options_name))

        # Периодически выполняйте коммит, чтобы не перегружать память
        if i % number_of_entries_for_the_commit == 0:
            connection.commit()
            print(f"Inserted {i} rows...")

    # Финальный коммит
    connection.commit()

def add_table_glass_types(num_rows):
    # Укажите имя таблицы и её колонки
    table_name = "glass_types"
    columns = "(idglass_types, name)"

    # Начинайте вставку данных
    for i in range(num_rows):
        idglass_types = i+1
        name = auto_glass_types[i]

        # Создайте SQL-запрос для вставки данных
        insert_query = f"""
        INSERT INTO {table_name} {columns} 
        VALUES (%s, %s)
        """
        cursor.execute(insert_query, (idglass_types, name))

        # Периодически выполняйте коммит, чтобы не перегружать память
        if i % number_of_entries_for_the_commit == 0:
            connection.commit()
            print(f"Inserted {i} rows...")

    # Финальный коммит
    connection.commit()
    
def add_table_glass_options(num_rows):
    # Укажите имя таблицы и её колонки
    table_name = "glass_options"
    columns = "(idglass_options, name, glass_type_name)"

    # Начинайте вставку данных
    for i in range(num_rows):
        idglass_options = i+1
        name = glass_options_list[i]
        glass_type_name = random.choice(auto_glass_types)

        # Создайте SQL-запрос для вставки данных
        insert_query = f"""
        INSERT INTO {table_name} {columns} 
        VALUES (%s, %s, %s)
        """
        cursor.execute(insert_query, (idglass_options, name, glass_type_name))

        # Периодически выполняйте коммит, чтобы не перегружать память
        if i % number_of_entries_for_the_commit == 0:
            connection.commit()
            print(f"Inserted {i} rows...")

    # Финальный коммит
    connection.commit()  

def add_table_brands(num_rows):
    # Укажите имя таблицы и её колонки
    table_name = "brands"
    columns = "(idbrands, name)"

    # Начинайте вставку данных
    for i in range(num_rows):
        idbrands = i+1
        name = car_brands[i]

        # Создайте SQL-запрос для вставки данных
        insert_query = f"""
        INSERT INTO {table_name} {columns} 
        VALUES (%s, %s)
        """
        cursor.execute(insert_query, (idbrands, name))

        # Периодически выполняйте коммит, чтобы не перегружать память
        if i % number_of_entries_for_the_commit == 0:
            connection.commit()
            print(f"Inserted {i} rows...")

    # Финальный коммит
    connection.commit()

def add_table_models(num_rows):
    # Укажите имя таблицы и её колонки
    table_name = "models"
    columns = "(idmodels, name, brand_name)"

    # Начинайте вставку данных
    for i in range(num_rows):
        idmodels = i+1
        name = car_models[i]
        brand_name = car_brands[i]

        # Создайте SQL-запрос для вставки данных
        insert_query = f"""
        INSERT INTO {table_name} {columns} 
        VALUES (%s, %s, %s)
        """
        cursor.execute(insert_query, (idmodels, name, brand_name))

        # Периодически выполняйте коммит, чтобы не перегружать память
        if i % number_of_entries_for_the_commit == 0:
            connection.commit()
            print(f"Inserted {i} rows...")

    # Финальный коммит
    connection.commit()  

def add_table_year_model(num_rows):
    # Укажите имя таблицы и её колонки
    table_name = "year_model"
    columns = "(idyear_model, name, model_name)"

    # Начинайте вставку данных
    for i in range(num_rows):
        idyear_model = i+1
        name = year_ranges[i]
        model_name = car_models[i]

        # Создайте SQL-запрос для вставки данных
        insert_query = f"""
        INSERT INTO {table_name} {columns} 
        VALUES (%s, %s, %s)
        """
        cursor.execute(insert_query, (idyear_model, name, model_name))

        # Периодически выполняйте коммит, чтобы не перегружать память
        if i % number_of_entries_for_the_commit == 0:
            connection.commit()
            print(f"Inserted {i} rows...")

    # Финальный коммит
    connection.commit() 

start_time = time.time()
#add_table_brands(num_rows_brands)
#add_table_models(num_rows_models)
#add_table_year_model(num_rows_year_model)
#add_table_glass_types(num_rows_typeproducts)
#add_table_glass_options(num_rows_optionsproducts)
add_table_products(num_rows_products)



# Закройте соединение
cursor.close()
connection.close()

end_time = time.time()
print(f"{end_time - start_time} seconds.")