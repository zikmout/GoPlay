type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip_code"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Simon")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "ashish", City: "New Delhi", Zipcode: "110075"},
		{Name: "simon", City: "Paris", Zipcode: "75015"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}