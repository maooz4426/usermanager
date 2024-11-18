import { makeApi, Zodios, type ZodiosOptions } from "@zodios/core";
import { z } from "zod";

const User = z
  .object({ uuid: z.string().uuid(), name: z.string(), email: z.string() })
  .partial()
  .passthrough();
const UserSignupRequest = z
  .object({ name: z.string(), email: z.string(), password: z.string() })
  .passthrough();
const UserSignupResponse = z.object({ uuid: z.string().uuid() }).passthrough();
const UserSignInRequest = z
  .object({ email: z.string(), password: z.string() })
  .passthrough();
const UserSignInResponse = z.object({ uuid: z.string().uuid() }).passthrough();
const UserSignInError = z
  .object({ code: z.string(), errors: z.array(z.string()) })
  .passthrough();

export const schemas = {
  User,
  UserSignupRequest,
  UserSignupResponse,
  UserSignInRequest,
  UserSignInResponse,
  UserSignInError,
};

const endpoints = makeApi([
  {
    method: "get",
    path: "/users/:uuid",
    alias: "getUser",
    description: `user取得`,
    requestFormat: "json",
    response: User,
  },
  {
    method: "post",
    path: "/users/signin",
    alias: "UserSignIn",
    description: `userのsigninリクエスト`,
    requestFormat: "json",
    parameters: [
      {
        name: "body",
        type: "Body",
        schema: z
          .object({ email: z.string(), password: z.string() })
          .passthrough(),
      },
    ],
    response: z.object({ uuid: z.string().uuid() }).passthrough(),
    errors: [
      {
        status: 404,
        schema: UserSignInError,
      },
    ],
  },
  {
    method: "post",
    path: "/users/signup",
    alias: "UserSignUp",
    description: `userのsignup`,
    requestFormat: "json",
    parameters: [
      {
        name: "body",
        type: "Body",
        schema: z
          .object({ name: z.string(), email: z.string(), password: z.string() })
          .passthrough(),
      },
    ],
    response: z.object({ uuid: z.string().uuid() }).passthrough(),
  },
]);

export const usermanageAPI = new Zodios(endpoints);

