import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { Observable } from "rxjs";
import { environment } from "../environments/environment";

@Injectable({
  providedIn: "root",
})
export class ApiService {
  private baseUrl = environment.baseUrl;
  private apiUrl = `${this.baseUrl}/pods`;

  constructor(private http: HttpClient) {}

  getData(): Observable<any> {
    console.log(this.apiUrl);
    return this.http.get<any>(this.apiUrl);
  }
}
