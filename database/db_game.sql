-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 06, 2022 at 03:55 PM
-- Server version: 10.4.22-MariaDB
-- PHP Version: 8.1.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_game`
--

-- --------------------------------------------------------

--
-- Table structure for table `age_rating_categories`
--

CREATE TABLE `age_rating_categories` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` longtext DEFAULT NULL,
  `description` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `age_rating_categories`
--

INSERT INTO `age_rating_categories` (`id`, `name`, `description`, `created_at`, `updated_at`) VALUES
(1, 'RP', 'Rating Pending', '2022-03-06 21:16:07.106', '2022-03-06 21:16:07.106'),
(2, 'EC', 'Early Childhood', '2022-03-06 21:16:18.625', '2022-03-06 21:16:18.625'),
(3, 'E', 'Everyone', '2022-03-06 21:16:31.783', '2022-03-06 21:16:31.783'),
(4, 'T', 'Teen', '2022-03-06 21:16:42.363', '2022-03-06 21:16:42.363'),
(5, 'M', 'Mature', '2022-03-06 21:16:51.967', '2022-03-06 21:16:51.967'),
(6, 'A', 'Adult', '2022-03-06 21:17:00.263', '2022-03-06 21:17:00.263'),
(7, 'PG', 'Parental Guidance', '2022-03-06 21:17:30.722', '2022-03-06 21:17:30.722');

-- --------------------------------------------------------

--
-- Table structure for table `developers`
--

CREATE TABLE `developers` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `developers`
--

INSERT INTO `developers` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Trophy Games', '2022-03-06 21:08:28.807', '2022-03-06 21:08:28.807'),
(2, 'FromSoftware Inc', '2022-03-06 21:09:01.338', '2022-03-06 21:09:01.338'),
(3, 'Dread Hunger Team', '2022-03-06 21:10:01.530', '2022-03-06 21:10:01.530'),
(4, 'Rockstar North', '2022-03-06 21:10:57.436', '2022-03-06 21:10:57.436'),
(5, 'Bungie', '2022-03-06 21:11:13.828', '2022-03-06 21:11:13.828'),
(6, 'Valve', '2022-03-06 21:11:41.700', '2022-03-06 21:11:41.700'),
(7, 'Visual Concepts', '2022-03-06 21:12:00.985', '2022-03-06 21:12:00.985'),
(8, 'Facepunch Studios', '2022-03-06 21:12:42.085', '2022-03-06 21:12:42.085'),
(9, 'Steel Wool Studios', '2022-03-06 21:13:16.275', '2022-03-06 21:13:16.275'),
(10, 'Sports Interactive', '2022-03-06 21:14:01.129', '2022-03-06 21:14:01.129');

-- --------------------------------------------------------

--
-- Table structure for table `games`
--

CREATE TABLE `games` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `title` longtext DEFAULT NULL,
  `about` longtext DEFAULT NULL,
  `year` bigint(20) DEFAULT NULL,
  `developer_id` bigint(20) UNSIGNED DEFAULT NULL,
  `publisher_id` bigint(20) UNSIGNED DEFAULT NULL,
  `genre_id` bigint(20) UNSIGNED DEFAULT NULL,
  `age_rating_category_id` bigint(20) UNSIGNED DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `games`
--

INSERT INTO `games` (`id`, `title`, `about`, `year`, `developer_id`, `publisher_id`, `genre_id`, `age_rating_category_id`, `created_at`, `updated_at`) VALUES
(1, 'ELDEN RING', 'THE NEW FANTASY ACTION RPG. Rise, Tarnished, and be guided by grace to brandish the power of the Elden Ring and become an Elden Lord in the Lands Between.', 2022, 2, 2, 1, 4, '2022-03-06 21:28:34.080', '2022-03-06 21:28:34.080'),
(2, 'DREAD HUNGER', 'A game of survival and betrayal. Eight Explorers path their ship through the unforgiving Arctic. Among the crew, two traitors call on dark powers to undermine them.', 2022, 3, 3, 1, 4, '2022-03-06 21:29:56.052', '2022-03-06 21:29:56.052'),
(3, 'Grand Theft Auto V', 'When a young street hustler, a retired bank robber and a terrifying psychopath land themselves in trouble, they must pull off a series of dangerous heists to survive in a city in which they can trust nobody, least of all each other.', 2015, 4, 5, 2, 5, '2022-03-06 21:31:32.324', '2022-03-06 21:31:32.324');

-- --------------------------------------------------------

--
-- Table structure for table `genres`
--

CREATE TABLE `genres` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `genres`
--

INSERT INTO `genres` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Action', '2022-03-06 21:06:02.497', '2022-03-06 21:06:02.497'),
(2, 'Adventure', '2022-03-06 21:06:09.856', '2022-03-06 21:06:09.856'),
(3, 'Casual', '2022-03-06 21:06:16.552', '2022-03-06 21:06:16.552'),
(4, 'Indie', '2022-03-06 21:06:22.617', '2022-03-06 21:06:22.617'),
(5, 'Massively Multiplayer', '2022-03-06 21:06:32.189', '2022-03-06 21:06:32.189'),
(6, 'Racing', '2022-03-06 21:06:38.649', '2022-03-06 21:06:38.649'),
(7, 'RPG', '2022-03-06 21:06:42.746', '2022-03-06 21:06:42.746'),
(8, 'Simulation', '2022-03-06 21:06:47.504', '2022-03-06 21:06:47.504'),
(9, 'Sports', '2022-03-06 21:06:51.788', '2022-03-06 21:06:51.788'),
(10, 'Strategy', '2022-03-06 21:06:55.914', '2022-03-06 21:06:55.914');

-- --------------------------------------------------------

--
-- Table structure for table `maximum_specifications`
--

CREATE TABLE `maximum_specifications` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `os` longtext DEFAULT NULL,
  `processor` longtext DEFAULT NULL,
  `memory` longtext DEFAULT NULL,
  `graphics` longtext DEFAULT NULL,
  `direct_x` longtext DEFAULT NULL,
  `network` longtext DEFAULT NULL,
  `storage` longtext DEFAULT NULL,
  `sound_card` longtext DEFAULT NULL,
  `game_id` bigint(20) UNSIGNED DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `maximum_specifications`
--

INSERT INTO `maximum_specifications` (`id`, `os`, `processor`, `memory`, `graphics`, `direct_x`, `network`, `storage`, `sound_card`, `game_id`, `created_at`, `updated_at`) VALUES
(1, 'Windows 10 64 Bit', 'Intel Core i5 3470 @3.2 GHz', '8 GB RAM', 'NVIDIA GTX 660 2GB', 'Directx 11', 'Broadband Network Connection', '72 GB Available Space', '100% DirectX 10 Compatible', 3, '2022-03-06 21:41:55.820', '2022-03-06 21:41:55.820'),
(2, 'Windows 10 64 Bit', 'INTEL CORE I7-8700K', '16 GB RAM', 'NVIDIA GEFORCE GTX 1070 8 GB', 'Version 12', 'Broadband Network Connection', '60 GB available space', 'Windows Compatible Audio Device', 1, '2022-03-06 21:44:57.234', '2022-03-06 21:44:57.234'),
(3, 'Windows 10 64 Bit', 'Intel Core i7-4770K', '8 GB RAM', 'NVIDIA GeForce 1060-6GB', 'Version 12', 'Broadband Network Connection', '15 GB available space', 'Windows Compatible Audio Device', 2, '2022-03-06 21:46:48.299', '2022-03-06 21:46:48.299');

-- --------------------------------------------------------

--
-- Table structure for table `minimum_specifications`
--

CREATE TABLE `minimum_specifications` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `os` longtext DEFAULT NULL,
  `processor` longtext DEFAULT NULL,
  `memory` longtext DEFAULT NULL,
  `graphics` longtext DEFAULT NULL,
  `direct_x` longtext DEFAULT NULL,
  `network` longtext DEFAULT NULL,
  `storage` longtext DEFAULT NULL,
  `sound_card` longtext DEFAULT NULL,
  `game_id` bigint(20) UNSIGNED DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `minimum_specifications`
--

INSERT INTO `minimum_specifications` (`id`, `os`, `processor`, `memory`, `graphics`, `direct_x`, `network`, `storage`, `sound_card`, `game_id`, `created_at`, `updated_at`) VALUES
(1, 'Windows 7 64 Bit Service Pack 1', 'Intel Core 2 Quad CPU Q6600 @ 2.40GHz (4 CPUs)', '4 GB RAM', 'NVIDIA 9800 GT 1GB', 'DirectX 10', 'Broadband Internet Connection', '72 GB available space', '100% DirectX 10 compatible', 3, '2022-03-06 21:39:29.335', '2022-03-06 21:43:10.217'),
(2, 'Windows 10', 'INTEL CORE I5-8400', '12 GB RAM', 'NVIDIA GEFORCE GTX 1060 3 GB', 'Version 12', 'Broadband Internet Connection', '60 GB available space', 'Windows Compatible Audio Device', 1, '2022-03-06 21:45:51.101', '2022-03-06 21:45:51.101'),
(3, 'Windows 7', 'Quad-core Intel or AMD, 2.5 GHz', '8 GB RAM', 'NVIDIA GEFORCE GTX 1060 3 GB', 'Version 12', 'Broadband Internet Connection', '15 GB available space', 'Windows Compatible Audio Device', 2, '2022-03-06 21:47:23.887', '2022-03-06 21:47:23.887');

-- --------------------------------------------------------

--
-- Table structure for table `publishers`
--

CREATE TABLE `publishers` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `publishers`
--

INSERT INTO `publishers` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Trophy Games', '2022-03-06 21:08:13.091', '2022-03-06 21:08:13.091'),
(2, 'BANDAI Namco', '2022-03-06 21:09:23.968', '2022-03-06 21:09:23.968'),
(3, 'Digital Confectioners', '2022-03-06 21:09:49.998', '2022-03-06 21:09:49.998'),
(4, 'Techland', '2022-03-06 21:10:27.766', '2022-03-06 21:10:27.766'),
(5, 'Rockstar Games', '2022-03-06 21:10:49.356', '2022-03-06 21:10:49.356'),
(6, 'Bungie', '2022-03-06 21:11:20.755', '2022-03-06 21:11:20.755'),
(7, 'Valve', '2022-03-06 21:11:36.767', '2022-03-06 21:11:36.767'),
(8, '2K', '2022-03-06 21:12:07.393', '2022-03-06 21:12:07.393'),
(9, 'Facepunch Studios', '2022-03-06 21:12:29.844', '2022-03-06 21:12:29.844'),
(10, 'ScottGames', '2022-03-06 21:13:36.630', '2022-03-06 21:13:36.630'),
(11, 'SEGA', '2022-03-06 21:13:48.559', '2022-03-06 21:13:48.559');

-- --------------------------------------------------------

--
-- Table structure for table `ratings`
--

CREATE TABLE `ratings` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `rating` bigint(20) DEFAULT NULL,
  `game_id` bigint(20) UNSIGNED DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `ratings`
--

INSERT INTO `ratings` (`id`, `rating`, `game_id`, `created_at`, `updated_at`) VALUES
(1, 70, 1, '2022-03-06 21:52:33.133', '2022-03-06 21:52:33.133'),
(2, 75, 1, '2022-03-06 21:52:37.536', '2022-03-06 21:52:37.536'),
(3, 65, 1, '2022-03-06 21:52:40.209', '2022-03-06 21:52:40.209'),
(4, 30, 2, '2022-03-06 21:52:55.623', '2022-03-06 21:52:55.623'),
(5, 50, 2, '2022-03-06 21:53:00.385', '2022-03-06 21:53:00.385'),
(6, 55, 2, '2022-03-06 21:53:03.210', '2022-03-06 21:53:03.210'),
(7, 89, 3, '2022-03-06 21:53:09.644', '2022-03-06 21:53:09.644'),
(8, 90, 3, '2022-03-06 21:53:12.910', '2022-03-06 21:53:12.910'),
(9, 78, 3, '2022-03-06 21:53:15.882', '2022-03-06 21:53:15.882'),
(10, 84, 3, '2022-03-06 21:53:18.841', '2022-03-06 21:53:18.841');

-- --------------------------------------------------------

--
-- Table structure for table `reviews`
--

CREATE TABLE `reviews` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `comment` longtext DEFAULT NULL,
  `game_id` bigint(20) UNSIGNED DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `reviews`
--

INSERT INTO `reviews` (`id`, `comment`, `game_id`, `created_at`, `updated_at`) VALUES
(1, 'fun but bad player base very angry very rude but fun game when have good people', 2, '2022-03-06 21:48:19.904', '2022-03-06 21:48:19.904'),
(2, 'Great if you wanna learn Chinese.', 2, '2022-03-06 21:49:07.440', '2022-03-06 21:49:07.440'),
(3, 'It\'s a great game, you dont know who to trust.', 2, '2022-03-06 21:50:09.054', '2022-03-06 21:50:09.054'),
(4, 'I would write a more indepth review, but im too busy playing Elden Ring.', 1, '2022-03-06 21:50:42.892', '2022-03-06 21:50:42.892'),
(5, 'A near perfect game that needs some performance improvements', 1, '2022-03-06 21:50:56.594', '2022-03-06 21:50:56.594'),
(6, 'Every time I hop on this game i either have a great time or am contemplating suicide. Best game ive ever played.', 1, '2022-03-06 21:51:05.276', '2022-03-06 21:51:05.276'),
(7, 'Highly recommend for young children to learn manners, respect for others, sharing, and to be a good citizen', 3, '2022-03-06 21:51:32.605', '2022-03-06 21:51:32.605'),
(8, 'Players are toxic or cheaters! a clean nice player is very very veeeeeeeeeeeeeeeery rare >.<', 3, '2022-03-06 21:51:44.993', '2022-03-06 21:51:44.993'),
(9, 'booted into a lobby and saw my address in the chat, this game makes me feel safe', 3, '2022-03-06 21:52:02.830', '2022-03-06 21:52:02.830');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `username` varchar(191) NOT NULL,
  `email` varchar(191) NOT NULL,
  `password` longtext NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `email`, `password`, `created_at`, `updated_at`) VALUES
(1, 'cahayoyo', 'cahayoyo@gmail.com', '$2a$10$W6Gf42f.ty/P61g/eiinLOnDdYBpyHdDq5M6el82B/S8r5Wfk/mi2', '2022-03-06 21:04:46.473', '2022-03-06 21:04:46.473');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `age_rating_categories`
--
ALTER TABLE `age_rating_categories`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `developers`
--
ALTER TABLE `developers`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `games`
--
ALTER TABLE `games`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_developers_game` (`developer_id`),
  ADD KEY `fk_publishers_game` (`publisher_id`),
  ADD KEY `fk_genres_game` (`genre_id`),
  ADD KEY `fk_age_rating_categories_game` (`age_rating_category_id`);

--
-- Indexes for table `genres`
--
ALTER TABLE `genres`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `maximum_specifications`
--
ALTER TABLE `maximum_specifications`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_games_maximum_specification` (`game_id`);

--
-- Indexes for table `minimum_specifications`
--
ALTER TABLE `minimum_specifications`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_games_minimum_specification` (`game_id`);

--
-- Indexes for table `publishers`
--
ALTER TABLE `publishers`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `ratings`
--
ALTER TABLE `ratings`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_games_rating` (`game_id`);

--
-- Indexes for table `reviews`
--
ALTER TABLE `reviews`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_games_review` (`game_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`),
  ADD UNIQUE KEY `email` (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `age_rating_categories`
--
ALTER TABLE `age_rating_categories`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `developers`
--
ALTER TABLE `developers`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `games`
--
ALTER TABLE `games`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `genres`
--
ALTER TABLE `genres`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `maximum_specifications`
--
ALTER TABLE `maximum_specifications`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `minimum_specifications`
--
ALTER TABLE `minimum_specifications`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `publishers`
--
ALTER TABLE `publishers`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `ratings`
--
ALTER TABLE `ratings`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `reviews`
--
ALTER TABLE `reviews`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `games`
--
ALTER TABLE `games`
  ADD CONSTRAINT `fk_age_rating_categories_game` FOREIGN KEY (`age_rating_category_id`) REFERENCES `age_rating_categories` (`id`),
  ADD CONSTRAINT `fk_developers_game` FOREIGN KEY (`developer_id`) REFERENCES `developers` (`id`),
  ADD CONSTRAINT `fk_genres_game` FOREIGN KEY (`genre_id`) REFERENCES `genres` (`id`),
  ADD CONSTRAINT `fk_publishers_game` FOREIGN KEY (`publisher_id`) REFERENCES `publishers` (`id`);

--
-- Constraints for table `maximum_specifications`
--
ALTER TABLE `maximum_specifications`
  ADD CONSTRAINT `fk_games_maximum_specification` FOREIGN KEY (`game_id`) REFERENCES `games` (`id`);

--
-- Constraints for table `minimum_specifications`
--
ALTER TABLE `minimum_specifications`
  ADD CONSTRAINT `fk_games_minimum_specification` FOREIGN KEY (`game_id`) REFERENCES `games` (`id`);

--
-- Constraints for table `ratings`
--
ALTER TABLE `ratings`
  ADD CONSTRAINT `fk_games_rating` FOREIGN KEY (`game_id`) REFERENCES `games` (`id`);

--
-- Constraints for table `reviews`
--
ALTER TABLE `reviews`
  ADD CONSTRAINT `fk_games_review` FOREIGN KEY (`game_id`) REFERENCES `games` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
