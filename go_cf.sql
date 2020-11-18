-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Nov 18, 2020 at 09:02 PM
-- Server version: 10.4.11-MariaDB
-- PHP Version: 7.4.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `go_cf`
--

-- --------------------------------------------------------

--
-- Table structure for table `spk_anemia_gejala_penyakit`
--

CREATE TABLE `spk_anemia_gejala_penyakit` (
  `id` int(11) NOT NULL,
  `Kode_gejala` varchar(5) NOT NULL,
  `Nama_gejala` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `spk_anemia_gejala_penyakit`
--

INSERT INTO `spk_anemia_gejala_penyakit` (`id`, `Kode_gejala`, `Nama_gejala`) VALUES
(1, 'G01', 'Terasa lemas diseluruh tubuh'),
(2, 'G02', 'Merasakan sakit kepala'),
(3, 'G03', 'Nyeri pada dada'),
(4, 'G04', 'Demam'),
(5, 'G05', 'Keluar darah dari hidung/Mimisan'),
(6, 'G06', 'Kaki dan tangan terasa dingin'),
(7, 'G07', 'Kesemutan pada kaki'),
(8, 'G08', 'Kulit tampak pucat'),
(9, 'G09', 'Merasakan muntah-muntah/mual'),
(10, 'G10', 'Nyeri pada panggul hingga ke paha'),
(11, 'G11', 'Nyeri pada ulu hati'),
(12, 'G12', 'BAB mengeluarkan darah');

-- --------------------------------------------------------

--
-- Table structure for table `spk_anemia_penyakit`
--

CREATE TABLE `spk_anemia_penyakit` (
  `id` int(11) NOT NULL,
  `Kode_penyakit` varchar(5) NOT NULL,
  `Nama_penyakit` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `spk_anemia_penyakit`
--

INSERT INTO `spk_anemia_penyakit` (`id`, `Kode_penyakit`, `Nama_penyakit`) VALUES
(1, 'P01', 'Anemia Aplastik'),
(2, 'P02', 'Anemia Defisiensi Zat besi'),
(3, 'P03', 'Anemia Kremis/Kronik');

-- --------------------------------------------------------

--
-- Table structure for table `spk_anemia_rule`
--

CREATE TABLE `spk_anemia_rule` (
  `id` int(11) NOT NULL,
  `Bobotnya` double NOT NULL,
  `MB` double DEFAULT NULL,
  `MD` double DEFAULT NULL,
  `Gejalanya_id` int(11) DEFAULT NULL,
  `Penyakitnya_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `spk_anemia_rule`
--

INSERT INTO `spk_anemia_rule` (`id`, `Bobotnya`, `MB`, `MD`, `Gejalanya_id`, `Penyakitnya_id`) VALUES
(3, 0.3, 0.8, 0.2, 1, 1),
(4, 0.3, 0.8, 0.1, 2, 1),
(5, 0.4, 0.75, 0.2, 4, 1),
(6, 0.8, 0.7, 0.2, 5, 1),
(7, 0.4, 0.77, 0.2, 9, 1),
(8, 0.7, 0.8, 0.3, 11, 1),
(9, 0.3, 0.65, 0.2, 1, 2),
(10, 0.3, 0.7, 0.2, 2, 2),
(11, 0.8, 0.8, 0.2, 3, 2),
(12, 0.4, 0.65, 0.2, 6, 2),
(13, 0.5, 0.6, 0.2, 7, 2),
(14, 0.6, 0.65, 0.3, 8, 2),
(15, 0.5, 0.6, 0.3, 10, 2),
(16, 0.3, 0.6, 0.2, 1, 3),
(17, 0.3, 0.65, 0.2, 2, 3),
(18, 0.4, 0.7, 0.3, 4, 3),
(19, 0.4, 0.7, 0.25, 6, 3),
(20, 0.4, 0.7, 0.2, 9, 3),
(21, 0.5, 0.65, 0.1, 10, 3),
(22, 0.7, 0.8, 0.1, 11, 3),
(23, 0.8, 0.85, 0.1, 12, 3),
(24, 0.6, 0.65, 0.3, 8, 1);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `spk_anemia_gejala_penyakit`
--
ALTER TABLE `spk_anemia_gejala_penyakit`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `spk_anemia_penyakit`
--
ALTER TABLE `spk_anemia_penyakit`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `spk_anemia_rule`
--
ALTER TABLE `spk_anemia_rule`
  ADD PRIMARY KEY (`id`),
  ADD KEY `spk_anemia_rule_Gejalanya_id_9a4059ec_fk_spk_anemi` (`Gejalanya_id`),
  ADD KEY `spk_anemia_rule_Penyakitnya_id_a5b1ef12_fk_spk_anemi` (`Penyakitnya_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `spk_anemia_gejala_penyakit`
--
ALTER TABLE `spk_anemia_gejala_penyakit`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT for table `spk_anemia_penyakit`
--
ALTER TABLE `spk_anemia_penyakit`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `spk_anemia_rule`
--
ALTER TABLE `spk_anemia_rule`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=25;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `spk_anemia_rule`
--
ALTER TABLE `spk_anemia_rule`
  ADD CONSTRAINT `spk_anemia_rule_Gejalanya_id_9a4059ec_fk_spk_anemi` FOREIGN KEY (`Gejalanya_id`) REFERENCES `spk_anemia_gejala_penyakit` (`id`),
  ADD CONSTRAINT `spk_anemia_rule_Penyakitnya_id_a5b1ef12_fk_spk_anemi` FOREIGN KEY (`Penyakitnya_id`) REFERENCES `spk_anemia_penyakit` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
