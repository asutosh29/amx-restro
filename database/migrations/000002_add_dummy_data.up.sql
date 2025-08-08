INSERT INTO category (category_name) VALUES
('Appetizers'),
('Main Course'),
('Desserts'),
('Beverages'),
('Salads'),
('Soups'),
('Snacks'),
('Combos');

INSERT INTO items (category_id, item_name, item_description, img_url, price, isVeg, isAvailable) VALUES
(1, 'Spring Rolls', 'Crispy vegetarian rolls with spicy dip', '../public/images/items/default.jpg', 120.00, true, true),
(1, 'Chicken Wings', 'Spicy grilled chicken wings', '../public/images/items/default.jpg', 180.00, false, true),
(2, 'Paneer Butter Masala', 'Creamy paneer curry with rich tomato gravy', '../public/images/items/default.jpg', 250.00, true, true),
(2, 'Butter Chicken', 'Classic North Indian chicken curry', '../public/images/items/default.jpg', 280.00, false, true),
(3, 'Gulab Jamun', 'Soft, sweet syrup-soaked balls', '../public/images/items/default.jpg', 90.00, true, true),
(3, 'Chocolate Brownie', 'Fudgy brownie with chocolate syrup', '../public/images/items/default.jpg', 150.00, true, false),
(4, 'Mango Shake', 'Chilled mango shake with ice cream', '../public/images/items/default.jpg', 110.00, true, true),
(4, 'Cold Coffee', 'Iced coffee with whipped cream topping', '../public/images/items/default.jpg', 130.00, true, false),
(5, 'Greek Salad', 'Fresh salad with feta cheese, olives and cucumber', '../public/images/items/default.jpg', 160.00, true, true),
(5, 'Chicken Caesar Salad', 'Grilled chicken with romaine, parmesan, and croutons', '../public/images/items/default.jpg', 200.00, false, true),
(6, 'Tomato Soup', 'Classic tomato soup with herbs and cream', '../public/images/items/default.jpg', 100.00, true, true),
(6, 'Chicken Soup', 'Warm soup with shredded chicken and spices', '../public/images/items/default.jpg', 130.00, false, false),
(7, 'French Fries', 'Golden crispy fries with ketchup', '../public/images/items/default.jpg', 90.00, true, true),
(7, 'Samosa', 'Fried pastry with savory potato filling', '../public/images/items/default.jpg', 40.00, true, true),
(7, 'Egg Puff', 'Baked pastry with spiced egg filling', '../public/images/items/default.jpg', 60.00, false, true),
(8, 'Veg Thali Combo', 'Assorted vegetarian dishes with rice and roti', '../public/images/items/default.jpg', 220.00, true, true),
(8, 'Non-Veg Thali Combo', 'Chicken curry, rice, roti and salad', '../public/images/items/default.jpg', 280.00, false, true),
(8, 'Burger Combo', 'Burger with fries and drink', '../public/images/items/default.jpg', 200.00, false, true),
(4, 'Lassi', 'Sweet yogurt drink', '../public/images/items/default.jpg', 70.00, true, true),
(3, 'Rasgulla', 'Spongy sweet balls dipped in syrup', '../public/images/items/default.jpg', 85.00, true, true);

INSERT INTO tables(table_id, isAvailable) VALUES(1,1),(2,1), (3,1),(4,1), (5,1),(6,1);