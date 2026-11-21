# Generate two reproducible datasets with 100 observations each.
indexes <- 0:99

integer_values <- (indexes * 37) %% 100
float_values <- integer_values + 0.25

write.csv(
  data.frame(value = integer_values),
  "data/integers.csv",
  row.names = FALSE
)

write.csv(
  data.frame(value = float_values),
  "data/floats.csv",
  row.names = FALSE
)