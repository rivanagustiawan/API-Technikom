-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost:8889
-- Generation Time: Oct 11, 2022 at 07:04 AM
-- Server version: 5.7.34
-- PHP Version: 7.4.21

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `apicms`
--

-- --------------------------------------------------------

--
-- Table structure for table `deskripsi_pakets`
--

CREATE TABLE `deskripsi_pakets` (
  `id` int(11) NOT NULL,
  `id_paket` int(11) DEFAULT NULL,
  `deskripsi` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `deskripsi_pakets`
--

INSERT INTO `deskripsi_pakets` (`id`, `id_paket`, `deskripsi`, `created_at`, `updated_at`) VALUES
(3, 4, 'ini desk 1', '2022-10-09 19:24:38', '2022-10-09 19:24:38'),
(4, 4, 'ini deks 2', '2022-10-09 19:24:38', '2022-10-09 19:24:38'),
(5, 3, 'asdas', '2022-10-09 20:19:53', '2022-10-09 20:19:53'),
(7, 4, 'asdas', '2022-10-09 20:20:02', '2022-10-09 20:20:02'),
(8, 4, 'asdas', '2022-10-09 20:25:24', '2022-10-09 20:25:24'),
(10, 4, 'ini desk 4', '2022-10-10 10:18:55', '2022-10-10 10:18:55'),
(11, 4, 'ini desk 4', '2022-10-10 11:09:35', '2022-10-10 11:09:35');

-- --------------------------------------------------------

--
-- Table structure for table `kliens`
--

CREATE TABLE `kliens` (
  `id` int(10) UNSIGNED NOT NULL,
  `nama_klien` varchar(255) DEFAULT NULL,
  `pic_logo` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `kliens`
--

INSERT INTO `kliens` (`id`, `nama_klien`, `pic_logo`, `created_at`, `updated_at`) VALUES
(1, 'rivan', 'asfasdasdas', '2022-07-12 11:52:23', '2022-07-12 11:52:23'),
(2, 'riasd', 'imasda/asdafs.jpg', '2022-10-11 11:39:20', '2022-10-11 11:39:20');

-- --------------------------------------------------------

--
-- Table structure for table `konsuls`
--

CREATE TABLE `konsuls` (
  `id` int(10) UNSIGNED NOT NULL,
  `nama_lengkap` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `no_hp` varchar(255) DEFAULT NULL,
  `nama_perusahaan` varchar(255) DEFAULT NULL,
  `bidang_perusahaan` varchar(255) DEFAULT NULL,
  `kebutuhan` varchar(255) DEFAULT NULL,
  `deskripsi_kebutuhan` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `konsuls`
--

INSERT INTO `konsuls` (`id`, `nama_lengkap`, `email`, `no_hp`, `nama_perusahaan`, `bidang_perusahaan`, `kebutuhan`, `deskripsi_kebutuhan`, `created_at`, `updated_at`) VALUES
(1, '', '', '', '', '', '', '', '2022-07-15 05:29:45', '2022-07-15 05:29:45'),
(2, '', '', '', '', '', '', '', '2022-07-15 05:32:18', '2022-07-15 05:32:18'),
(3, '', '', '', '', '', '', '', '2022-07-15 05:33:10', '2022-07-15 05:33:10'),
(4, '', '', '', '', '', '', '', '2022-07-15 05:33:11', '2022-07-15 05:33:11'),
(5, '', '', '', '', '', '', '', '2022-07-15 05:33:54', '2022-07-15 05:33:54'),
(6, '', '', '', '', '', '', '', '2022-07-15 05:34:09', '2022-07-15 05:34:09'),
(7, '', '', '', '', '', '', '', '2022-07-17 21:37:34', '2022-07-17 21:37:34'),
(8, '', '', '', '', '', '', '', '2022-07-17 21:37:51', '2022-07-17 21:37:51'),
(9, '', '', '', '', '', '', '', '2022-07-21 17:37:58', '2022-07-21 17:37:58'),
(10, '', '', '', '', '', '', '', '2022-07-21 17:41:04', '2022-07-21 17:41:04'),
(11, '', '', '', '', '', '', '', '2022-07-21 19:38:19', '2022-07-21 19:38:19');

-- --------------------------------------------------------

--
-- Table structure for table `pakets`
--

CREATE TABLE `pakets` (
  `id` int(10) UNSIGNED NOT NULL,
  `nama_paket` varchar(255) DEFAULT NULL,
  `harga` varchar(255) DEFAULT NULL,
  `jenis_paket` varchar(255) DEFAULT NULL,
  `warna_label` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `pakets`
--

INSERT INTO `pakets` (`id`, `nama_paket`, `harga`, `jenis_paket`, `warna_label`, `created_at`, `updated_at`) VALUES
(4, 'paket4update', '400000', 'adasgasdasd', '1421412', '2022-10-09 18:49:27', '2022-10-09 23:05:39');

-- --------------------------------------------------------

--
-- Table structure for table `portofolios`
--

CREATE TABLE `portofolios` (
  `id` int(10) UNSIGNED NOT NULL,
  `id_klien` int(11) DEFAULT NULL,
  `nama_apps` varchar(255) DEFAULT NULL,
  `jenis_paket` varchar(255) DEFAULT NULL,
  `link_apps` varchar(255) DEFAULT NULL,
  `deskripsi_apps` varchar(255) DEFAULT NULL,
  `photo_apps1` varchar(255) DEFAULT NULL,
  `photo_apps2` varchar(255) DEFAULT NULL,
  `photo_apps3` varchar(255) DEFAULT NULL,
  `photo_apps4` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `portofolios`
--

INSERT INTO `portofolios` (`id`, `id_klien`, `nama_apps`, `jenis_paket`, `link_apps`, `deskripsi_apps`, `photo_apps1`, `photo_apps2`, `photo_apps3`, `photo_apps4`, `created_at`, `updated_at`) VALUES
(1, 1, 'tesporto', 'teuing', 'aasfas', 'asdas', '123', NULL, NULL, NULL, '2022-07-12 11:52:44', '2022-07-12 11:52:44');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `deskripsi_pakets`
--
ALTER TABLE `deskripsi_pakets`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `kliens`
--
ALTER TABLE `kliens`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `konsuls`
--
ALTER TABLE `konsuls`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `pakets`
--
ALTER TABLE `pakets`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `portofolios`
--
ALTER TABLE `portofolios`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `deskripsi_pakets`
--
ALTER TABLE `deskripsi_pakets`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `kliens`
--
ALTER TABLE `kliens`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `konsuls`
--
ALTER TABLE `konsuls`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `pakets`
--
ALTER TABLE `pakets`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `portofolios`
--
ALTER TABLE `portofolios`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
