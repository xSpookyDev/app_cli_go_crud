/**
 * Crear la tabla de clientes
 */

CREATE TABLE `clientes` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `nombre` varchar(100) NOT NULL,
    `correo` varchar(200) NOT NULL,
    `telefono` varchar(200) NOT NULL,
    `mensaje` varchar(200) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


/**
 * Insertar datos de ejemplo
 */
INSERT INTO `clientes` (`nombre`, `correo`, `telefono`, `mensaje`) VALUES
('John Doe', 'john.doe@example.com', '123-456-7890', 'Hello, this is John.'),
('Jane Smith', 'jane.smith@example.com', '098-765-4321', 'Hi, this is Jane.');