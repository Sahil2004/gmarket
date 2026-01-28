import { HttpEvent, HttpEventType, HttpHandlerFn, HttpRequest } from '@angular/common/http';
import { inject } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Observable, tap } from 'rxjs';
import { IError } from '../types';

export function errorHandlerInterceptor(
  req: HttpRequest<any>,
  next: HttpHandlerFn,
): Observable<HttpEvent<unknown>> {
  return next(req).pipe(
    tap((event) => {
      if (event.type === HttpEventType.Response) {
        if (event.status >= 400 && event.status < 500) {
          let _snackBar = inject(MatSnackBar);
          let snackBarRef = _snackBar.open((event.body as IError).message, 'Close', {
            duration: 3000,
          });
          snackBarRef.onAction().subscribe(() => {
            snackBarRef.dismiss();
          });
        }
      }
    }),
  );
}
