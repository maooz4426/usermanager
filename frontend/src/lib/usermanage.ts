import { makeApi, Zodios, type ZodiosOptions } from "@zodios/core";
import { z } from "zod";

const User = z
  .object({ uuid: z.string(), name: z.string(), email: z.string() })
  .partial()
  .passthrough();
const UserSignupRequest = z
  .object({ name: z.string(), email: z.string(), password: z.string() })
  .passthrough();
const UserSignupResponse = z.object({ uuid: z.string() }).passthrough();
const UserSignInRequest = z
  .object({ email: z.string(), password: z.string() })
  .passthrough();
const UserSignInResponse = z.object({ uuid: z.string() }).passthrough();
const UserSignInError = z.object({ errors: z.array(z.string()) }).passthrough();

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
    parameters: [
      {
        name: "uuid",
        type: "Path",
        schema: z.number().int(),
      },
    ],
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
        schema: UserSignInRequest,
      },
    ],
    response: z.object({ uuid: z.string() }).passthrough(),
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
        schema: UserSignupRequest,
      },
    ],
    response: z.object({ uuid: z.string() }).passthrough(),
  },
]);

export const usermanageAPI = new Zodios(endpoints);

export function createApiClient(baseUrl: string, options?: ZodiosOptions) {
  return new Zodios(baseUrl, endpoints, options);
}
