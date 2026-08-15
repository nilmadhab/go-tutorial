-- Create products table
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    category VARCHAR(100),
    stock INTEGER DEFAULT 0,
    image VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert sample data
INSERT INTO products (name, description, price, category, stock, image) VALUES
    ('Wireless Bluetooth Headphones', 'Premium noise-canceling headphones with 30-hour battery life', 149.99, 'Electronics', 50, 'headphones.jpg'),
    ('Cotton T-Shirt', 'Comfortable 100% cotton t-shirt available in multiple colors', 24.99, 'Clothing', 200, 'tshirt.jpg'),
    ('Stainless Steel Water Bottle', 'Insulated water bottle keeps drinks cold for 24 hours', 34.99, 'Home & Kitchen', 75, 'bottle.jpg'),
    ('Running Shoes', 'Lightweight running shoes with cushioned sole', 89.99, 'Sports', 120, 'shoes.jpg'),
    ('Laptop Stand', 'Adjustable aluminum laptop stand for better ergonomics', 49.99, 'Electronics', 85, 'stand.jpg')
ON CONFLICT DO NOTHING;
