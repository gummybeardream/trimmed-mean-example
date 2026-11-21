# Validate the shared datasets using R's base mean function.
integer_values <- read.csv("data/integers.csv")$value
float_values <- read.csv("data/floats.csv")$value

stopifnot(length(integer_values) == 100)
stopifnot(length(float_values) == 100)

trim_proportion <- 0.05

integer_trimmed_mean <- mean(integer_values, trim = trim_proportion)
float_trimmed_mean <- mean(float_values, trim = trim_proportion)

cat("R trimmed mean validation\n")
cat(sprintf(
  "Integer trimmed mean (5%%): %.10f\n",
  integer_trimmed_mean
))
cat(sprintf(
  "Floating-point trimmed mean (5%%): %.10f\n",
  float_trimmed_mean
))