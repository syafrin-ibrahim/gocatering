-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 02 Jul 2021 pada 16.45
-- Versi server: 10.4.13-MariaDB
-- Versi PHP: 7.4.7

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `gocatering`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `categories`
--

CREATE TABLE `categories` (
  `id` bigint(20) NOT NULL,
  `name` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `categories`
--

INSERT INTO `categories` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'weekly', '2021-07-02 07:01:22.960', '2021-07-02 07:37:06.629'),
(3, 'event', '2021-07-02 21:48:29.420', '2021-07-02 21:48:29.420'),
(4, 'healty', '2021-07-02 21:49:12.902', '2021-07-02 21:49:12.902'),
(5, 'event', '2021-07-02 22:11:00.071', '2021-07-02 22:11:00.071'),
(6, 'event-olahraga', '2021-07-02 22:12:27.174', '2021-07-02 22:12:27.174'),
(7, 'event-olahraga', '2021-07-02 22:13:01.356', '2021-07-02 22:13:01.356');

-- --------------------------------------------------------

--
-- Struktur dari tabel `images`
--

CREATE TABLE `images` (
  `id` bigint(20) NOT NULL,
  `paket_id` bigint(20) DEFAULT NULL,
  `file_name` longtext DEFAULT NULL,
  `is_main` tinyint(1) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Struktur dari tabel `pakets`
--

CREATE TABLE `pakets` (
  `id` bigint(20) NOT NULL,
  `category_id` bigint(20) DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `description` longtext DEFAULT NULL,
  `price` bigint(20) DEFAULT NULL,
  `discount` bigint(20) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `pakets`
--

INSERT INTO `pakets` (`id`, `category_id`, `name`, `description`, `price`, `discount`, `created_at`, `updated_at`) VALUES
(1, 1, 'Daily Lunch Box', 'Nasi putih, Ayam goreng/ ikan / udang,Sambal gr, kentang, ati, Tahu, tempe,bakwan,Mix vegetable,bakso,Sambal,Kerupuk', 35000, 0, '2021-07-02 09:11:33.623', '2021-07-02 09:11:33.623'),
(3, 3, 'Weight Loss', ' White Rice Garlic Butter Rice Seaweed Rice Mashed Potato Potato Wedges', 65000, 0, '2021-07-02 21:56:49.625', '2021-07-02 21:56:49.625'),
(6, 3, 'Wedding', 'Nasi Putih Nasi Goreng Ikan Bumbu Kuning Sambal Goreng Ati Soup/Capcay Buah Air Mineral', 45000, 0, '2021-07-02 22:00:11.051', '2021-07-02 22:00:11.051');

-- --------------------------------------------------------

--
-- Struktur dari tabel `regencies`
--

CREATE TABLE `regencies` (
  `id` bigint(20) NOT NULL,
  `name` longtext DEFAULT NULL,
  `shipping_cost` bigint(20) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `regencies`
--

INSERT INTO `regencies` (`id`, `name`, `shipping_cost`, `created_at`, `updated_at`) VALUES
(1, 'Kabupate Pohuwato', 120000, '2021-07-01 21:02:34.793', '2021-07-02 07:10:41.213'),
(8, 'Kabupaten Bone Bolango', 450000, '2021-07-02 21:49:43.820', '2021-07-02 21:49:43.820'),
(9, 'Kota Gorontalo', 150000, '2021-07-02 21:50:07.452', '2021-07-02 21:50:07.452');

-- --------------------------------------------------------

--
-- Struktur dari tabel `transactions`
--

CREATE TABLE `transactions` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) DEFAULT NULL,
  `paket_id` bigint(20) DEFAULT NULL,
  `quantity` bigint(20) DEFAULT NULL,
  `total` bigint(20) DEFAULT NULL,
  `location` longtext DEFAULT NULL,
  `regency_id` bigint(20) DEFAULT NULL,
  `status` longtext DEFAULT NULL,
  `payment_url` longtext DEFAULT NULL,
  `deliver_time` longtext DEFAULT NULL,
  `note` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `transactions`
--

INSERT INTO `transactions` (`id`, `user_id`, `paket_id`, `quantity`, `total`, `location`, `regency_id`, `status`, `payment_url`, `deliver_time`, `note`, `created_at`, `updated_at`) VALUES
(27, 3, 1, 15, 645000, 'jl sepati', 1, 'pending', 'https://app.sandbox.midtrans.com/snap/v2/vtweb/78a84ebd-d857-419e-b4f7-3d7165ebb8c3', '2021-06-29 22:57', '15 menit sebelum acara dimulai', '2021-07-02 20:07:24.851', '2021-07-02 20:07:25.364'),
(28, 3, 1, 25, 995000, 'jl sepati', 1, 'pending', 'https://app.sandbox.midtrans.com/snap/v2/vtweb/ad851ef5-5983-4eb7-a19d-0b4992ad6e7c', '2021-06-29 22:57', '15 menit sebelum acara dimulai', '2021-07-02 20:10:04.170', '2021-07-02 20:10:04.754'),
(30, 2, 3, 15, 1095000, 'jl ahmad yani', 1, 'pending', 'https://app.sandbox.midtrans.com/snap/v2/vtweb/52091dc0-4fe0-4a36-ae4f-6309b3439fe3', '2021-06-29 22:57', '15 menit sebelum acara dimulai', '2021-07-02 22:18:16.451', '2021-07-02 22:18:17.514');

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL,
  `full_name` longtext DEFAULT NULL,
  `mobile` longtext DEFAULT NULL,
  `email` longtext DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  `address` longtext DEFAULT NULL,
  `is_admin` tinyint(1) DEFAULT NULL,
  `image_url` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `full_name`, `mobile`, `email`, `password`, `address`, `is_admin`, `image_url`, `created_at`, `updated_at`) VALUES
(1, 'andika pratama', 'andika pratama', 'andika@gmail.com', '$2a$14$dNN1N2nlK78eTNo7s9y6ye3A65mP3u99zvycU88POajeskMl7NyHi', 'jl selayar', 1, '', '2021-07-02 11:14:25.353', '2021-07-02 11:14:25.353'),
(2, 'ana aluna', '1223344', 'ana@gmail.com', '$2a$14$/oVQp25q0eFdKoFVoXuN2OqbOPiSeTwOXphneiwqalANrw6uT/38e', 'tapa', 0, 'avatar/ana_aluna_avatar3.png', '2021-07-02 17:43:21.319', '2021-07-02 21:29:46.474'),
(3, 'gito sabara', '1223344', 'gito@gmail.com', '$2a$14$wWUyZTYa2u5pTY1z1Oa5jebZ9uNfswLcre/6forf4.yGxgG3KZrVK', 'tapa', 0, '', '2021-07-02 20:01:30.807', '2021-07-02 21:00:49.019'),
(4, 'dedi', 'dedi', 'dedi@gmail.com', '$2a$14$e43MSLQ4u97VhY/o6n2EJujYAtavEBFamxbCmSvzUE0M7fvvdkD5O', 'jl selangor', 0, '', '2021-07-02 22:14:41.832', '2021-07-02 22:14:41.832');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `categories`
--
ALTER TABLE `categories`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `images`
--
ALTER TABLE `images`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_pakets_image` (`paket_id`);

--
-- Indeks untuk tabel `pakets`
--
ALTER TABLE `pakets`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_categories_paket` (`category_id`);

--
-- Indeks untuk tabel `regencies`
--
ALTER TABLE `regencies`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_users_transaction` (`user_id`),
  ADD KEY `fk_pakets_transaction` (`paket_id`),
  ADD KEY `fk_transactions_regency` (`regency_id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `categories`
--
ALTER TABLE `categories`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT untuk tabel `images`
--
ALTER TABLE `images`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `pakets`
--
ALTER TABLE `pakets`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT untuk tabel `regencies`
--
ALTER TABLE `regencies`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT untuk tabel `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=31;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `images`
--
ALTER TABLE `images`
  ADD CONSTRAINT `fk_pakets_image` FOREIGN KEY (`paket_id`) REFERENCES `pakets` (`id`);

--
-- Ketidakleluasaan untuk tabel `pakets`
--
ALTER TABLE `pakets`
  ADD CONSTRAINT `fk_categories_paket` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`);

--
-- Ketidakleluasaan untuk tabel `transactions`
--
ALTER TABLE `transactions`
  ADD CONSTRAINT `fk_pakets_transaction` FOREIGN KEY (`paket_id`) REFERENCES `pakets` (`id`),
  ADD CONSTRAINT `fk_transactions_regency` FOREIGN KEY (`regency_id`) REFERENCES `regencies` (`id`),
  ADD CONSTRAINT `fk_users_transaction` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
