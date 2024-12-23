(def production-per-hour 221)

(defn production-rate
  "Returns the assembly line's production rate per hour,
   taking into account its success rate"
  [speed]
  (cond (< speed 1) (* 0.0 (* production-per-hour speed))
        (< speed 5) (* 1.0 (* production-per-hour speed))
        (< speed 9) (* 0.9 (* production-per-hour speed))
        (= speed 9) (* 0.8 (* production-per-hour speed))
        :else (* 0.77 (* production-per-hour speed))))

(defn working-items
  "Calculates how many working cars are produced per minute"
  [speed]
  (int (/ (production-rate speed) 60)))

(printf (str (working-items 6)))
