-- `todo-list`.tasks definition

CREATE TABLE `tasks` (
  `id` varchar(100) NOT NULL,
  `job` text NOT NULL,
  `is_completed` tinyint(1) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;