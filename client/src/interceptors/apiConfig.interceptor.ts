import { HttpEvent, HttpHandlerFn, HttpRequest } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../environments/environment';

export function apiConfigInterceptor(
  req: HttpRequest<any>,
  next: HttpHandlerFn,
): Observable<HttpEvent<unknown>> {
  const apiReq = req.clone({
    withCredentials: true,
    url: `${environment.API_BASE_URI}${req.url}`,
  });

  return next(apiReq);
}
