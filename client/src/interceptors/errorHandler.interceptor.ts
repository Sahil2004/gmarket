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
import { SKIP_TOAST } from '../contexts';

export function errorHandlerInterceptor(
  req: HttpRequest<any>,
  next: HttpHandlerFn,
): Observable<HttpEvent<unknown>> {
  let _snackBar = inject(MatSnackBar);
  const skipToast = req.context.get(SKIP_TOAST);
  return next(req).pipe(
    catchError((err: HttpErrorResponse) => {
      if (err.status >= 400 && err.status <= 500 && !skipToast) {
        let snackBarRef = _snackBar.open((err.error as IError).message, 'Close', {
          duration: 3000,
        });
        snackBarRef.onAction().subscribe(() => {
          snackBarRef.dismiss();
        });
      } else if (!skipToast) {
        let snackBarRef = _snackBar.open(err.message, 'Close', {
          duration: 3000,
        });
        snackBarRef.onAction().subscribe(() => {
          snackBarRef.dismiss();
        });
      }
      return throwError(() => err);
    }),
  );
}
