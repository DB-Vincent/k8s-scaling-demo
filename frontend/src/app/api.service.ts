import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { Observable } from "rxjs";
import { environment } from "../environments/environment";
import { ApiResponse } from "./interfaces/api.interfaces";

@Injectable({
  providedIn: "root",
})
export class ApiService {
  private baseUrl = environment.baseUrl;
  private apiUrl = `${this.baseUrl}/pods`;

  constructor(private http: HttpClient) {}

  getData(): Observable<ApiResponse> {
    return this.http.get<ApiResponse>(this.apiUrl);
  }
}
