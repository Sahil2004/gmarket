import {
  HttpErrorResponse,
  HttpEvent,
  HttpEventType,
  HttpHandlerFn,
  HttpRequest,
} from '@angular/common/http';
import { inject } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { catchError, EMPTY, Observable, throwError } from 'rxjs';
import { IError } from '../types';

export function errorHandlerInterceptor(
  req: HttpRequest<any>,
  next: HttpHandlerFn,
): Observable<HttpEvent<unknown>> {
  let _snackBar = inject(MatSnackBar);
  return next(req).pipe(
    catchError((err: HttpErrorResponse) => {
      if (err.status >= 400 && err.status <= 500) {
        let snackBarRef = _snackBar.open((err.error as IError).message, 'Close', {
          duration: 3000,
        });
        snackBarRef.onAction().subscribe(() => {
          snackBarRef.dismiss();
        });
        return EMPTY;
      }
      console.error('An unexpected error occurred:', err);
      return throwError(() => err);
    }),
  );
}
