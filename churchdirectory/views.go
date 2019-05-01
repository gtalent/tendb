/*
   Copyright 2017 gtalent2@gmail.com

   This Source Code Form is subject to the terms of the Mozilla Public
   License, v. 2.0. If a copy of the MPL was not distributed with this
   file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package churchdirectory

import (
	"encoding/json"
	"net/http"

	"github.com/gtalent/tendb"
)

func SetupViews(prefix string, s *http.Server) {
	if len(prefix) > 0 && prefix[len(prefix)-1] != '/' {
		prefix += "/"
	}
	//s.HandleFunc(prefix+"directory_page", directoryPage)
}

func directoryPage(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	type request struct {
		MembershipStatus string `json:"membership_status"`
		Start            int    `json:"start"`
		End              int    `json:"end"`
	}
	var rqst request
	err := json.NewDecoder(r.Body).Decode(&rqst)
	if err != nil {
		return
	}
	//ms := rqst.MembershipStatus
	var eventBaptism EventType
	db, err := tendb.OpenDatabase()
	if err != nil {
		return
	}
	db.Where("member = ?", EventBaptism).First(&eventBaptism)
	//	filt = _build_membership_filter(ms)
	//	people := Person.objects.filter(filt).order_by('last_name', 'first_name')
	//	try:
	//	    start = data['start']
	//	except KeyError:
	//	    start = 0
	//	try:
	//	    end = data['end']
	//	except KeyError:
	//	    end = len(people)
	//	people = people[start:end]
	//var out []Person
	//for _, p := range people {
	//	out = append(out, _jsonify_person(p))
	//}
	//return HttpResponse(json.dumps(out, indent=3), content_type='application/json')
}
