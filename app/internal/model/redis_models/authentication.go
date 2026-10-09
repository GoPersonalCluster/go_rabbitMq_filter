package redismodels

import (
	"time"
)

type AuthenticationIP struct {
	Score     float64   `json:"score"`
	CreatedAt time.Time `json:"createdat"`
}

func NewAuthenticationIp() *AuthenticationIP {
	return &AuthenticationIP{
		Score:     100,
		CreatedAt: time.Now(),
	}
}

type AuthenticationAccount struct {
	UserId    uint      `json:"userid"`
	Score     float64   `json:"score"`
	CreatedAt time.Time `json:"createdat"`
}

type AuthenticationDevice struct {
	Device string `json:"device"`
	UserId uint   `json:"userid"`
}

/*
--0 Validar se a conta de usuário e IP são compatíveis, se o usuário pertence a um grupo de IP

--1 Ataques de spam de um grupo de usuários devem ser reconhecidos e reduzir o score
do IP deste grupo, a conta e IP devem ser bloqueadas caso haja indicios de spam

--1.1 Deve haver uma classificação de riscos para a tomada de decisão de
bloqueio para um grupo de usuários com notificação antecipada
-- Gravidade do ataque
--1 uma tentativa de autenticação com IP ou usuário comum que tentou autenticar próximo
ao limite de tempo do reset multiplicado de sua quantidade limite de chances
para autenticar antes de levar um bloqueio

--1.2 identificar ataques de spam para grupos de usuários, deve verificar se o
dispositivo que está acessando é comum para várias contas



--2 Deve haver um mecanismo que possa identificar multiplos dispositivos
que acessam a mesma conta e multiplas contas que acessam o mesmo dispotivo

--3 Deve haver um mecanismo de duplo fator de autenticação que solicite a autenticação
após a identificação de multiplas contas de usuário acessando o mesmo dispositivo

--4 a falha de autenticação deve diminuir o score do IP quando o IP não pertencer
a um grupo de usuários
--5 falhas de autenticação devem reduzir o score da conta ou IP


--6 deve haver um mecanismo de armazenamento que salve tentativas de acesso
de uma conta em multiplos dispositivos, um dispositivo com acesso de multiplas contas,
dispositivos que foram identificados em um ataque de spam


--F O Score deve ser superior á 0 para autenticar
--F O score deve possuir um tempo pré definido para ser reiniciado
--F renovação do token deve ser feita apenas pelo mesmo dispositivo que gerou o token
--F um IP que não esteja dentro de um grupo de usuários não pode acessar um
IP que esteja reservado













*/
