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
-- O Score deve ser superior á 0 para autenticar
-- O score deve possuir um tempo pré definido para ser reiniciado
-- A renovação do token deve ser feita apenas pelo mesmo dispositivo que gerou o token

-- falhas de autenticação devem reduzir o score da conta ou IP
-- a falha de autenticação deve diminuir o score do IP quando o IP não pertencer
a um grupo de usuários

-- Ataques de spam de um grupo de usuários devem ser reconhecidos e reduzir o score
do IP deste grupo

-- Deve haver uma classificação de riscos para a tomada de decisão de
bloqueio para um grupo de usuários com notificação antecipada
-- Gravidade do ataque

-- Deve haver um mecanismo que possa identificar multiplos dispositivos
que acessam a mesma conta e multiplas contas que acessam o mesmo dispotivo

-- Deve haver um mecanismo de duplo fator de autenticação que solicite a autenticação
após a identificação de multiplas contas de usuário acessando o mesmo dispositivo

-- deve haver um mecanismo de armazenamento que salve tentativas de acesso
de uma conta em multiplos dispositivos, um dispositivo com acesso de multiplas contas,
dispositivos que foram identificados em um ataque de spam




*/
